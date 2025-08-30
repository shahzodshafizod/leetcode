package graphs

/*
Problem:
There are a total of n courses to take, labeled from 0 to n-1.
example: n=6; 6 courses: 0, 1, 2, 3, 4, 5

Some courses have prerequisite courses. This is expressed as a pair i.e. [1,0]
which indicates you must take course 0 before taking course 1.
	that's: (0) -> (1)

Given the total number of courses and an array of prerequisite pairs,
return if it is possible to finish all courses.
	- n=6, prerequisites: [[1,0],[2,1],[2,5],[0,3],[4,3],[3,5],[4,5]]

	 +---(3)--->(4)			-> 	Is it cyclic?
	 |    ^      ^				Is there a cycle?
	 v    |      |				Can we find a cycle within this graph?
	(0)  (5)-----+
	 |           |			->	No, there is no any cycle. That's, it is possible to finish all courses.
	 |           v
	 +-->(1)--->(2)			->	True

There is a cycle, if you change the order in the third prerequisite: [2,5] -> [5,2]

	 +---(3)--->(4)			-> 	Is it cyclic?
	 |    ^      ^				Is there a cycle?
	 v    |      |				Can we find a cycle within this graph?
	(0)  (5)<----+
	 |           |			->	Yes, there is a cycle: 3-0-1-2-5. That's, it's not possible to finish all courses.
	 |           |
	 +-->(1)--->(2)			->	False

Step 1: Verify the constraints
	- Can we have courses unconnected to other courses?
		: Yes, account for separate course chains.
	        (0)<---(3)    (4)--->(7)
	         |             ^      |
	         v             |      |
	        (1)--->(2)    (5)<----+
	- Constraints:
		- 1 <= numCourses <= 2000
		- 0 <= prerequisites.length <= 5000
		- prerequisites[i].length == 2
		- 0 <= ai, bi < numCourses
		- All the pairs prerequisites[i] are unique.
Step 2: Write out some test cases

Topological Sort:
	- Every vertex has an indegree factor
	- Graph needs to have edges that are directed
	- Indegree is represented as "how many connections are coming into the vertex"
	       (3)
	        |
	        v
	(2)--->(1)--->(4)
	        ^
	        |
	       (5)			Indegree: 3 (how many edges are pointing into the vertex)

Directed Acyclic Graph

*/

// https://leetcode.com/problems/course-schedule/

// Topological Sort via BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make([][]int, numCourses)

	inDegree := make([]int, numCourses)
	for _, pair := range prerequisites {
		inDegree[pair[0]]++

		if adjList[pair[1]] == nil {
			adjList[pair[1]] = []int{pair[0]}
		} else {
			adjList[pair[1]] = append(adjList[pair[1]], pair[0])
		}
	}

	stack := make([]int, 0)

	for i := range numCourses {
		if inDegree[i] == 0 {
			stack = append(stack, i)
		}
	}

	count := 0
	for length := len(stack); length > 0; length = len(stack) {
		count++

		adjacent := adjList[stack[0]]
		for i, listLength := 0, len(adjacent); i < listLength; i++ {
			inDegree[adjacent[i]]--
			if inDegree[adjacent[i]] == 0 {
				stack = append(stack, adjacent[i])
			}
		}

		stack = stack[1:]
	}

	return count == numCourses
}

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	var adjList = make([][]int, numCourses)
// 	for _, pair := range prerequisites {
// 		if adjList[pair[1]] == nil {
// 			adjList[pair[1]] = []int{pair[0]}
// 		} else {
// 			adjList[pair[1]] = append(adjList[pair[1]], pair[0])
// 		}
// 	}
// 	for i := 0; i < numCourses; i++ {
// 		if adjList[i] != nil {
// 			var queue = append([]int{}, adjList[i]...)
// 			var seen = make(map[int]bool)
// 			for queueLen := len(queue); queueLen > 0; queueLen = len(queue) {
// 				for qi := 0; qi < queueLen; qi++ {
// 					if adjList[queue[qi]] == nil || seen[queue[qi]] {
// 						continue
// 					}
// 					seen[queue[qi]] = true
// 					if &adjList[queue[qi]] == &adjList[i] {
// 						return false
// 					}
// 					queue = append(queue, adjList[queue[qi]]...)
// 				}
// 				queue = queue[queueLen:]
// 			}
// 		}
// 	}
// 	return true
// }
