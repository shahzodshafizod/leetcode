package graphs

import "container/list"

// https://leetcode.com/problems/course-schedule-iv/

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adjList := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	var src, dst int
	for idx := range prerequisites {
		src, dst = prerequisites[idx][0], prerequisites[idx][1]
		adjList[src] = append(adjList[src], dst)
		indegree[dst]++
	}
	queue := list.New()
	reach := make([][]bool, numCourses)
	for course := 0; course < numCourses; course++ {
		reach[course] = make([]bool, numCourses)
		if indegree[course] == 0 {
			queue.PushBack(course)
		}
	}
	for queue.Len() > 0 {
		src = queue.Remove(queue.Front()).(int)
		for _, dst := range adjList[src] {
			reach[dst][src] = true
			for node := 0; node < numCourses; node++ {
				reach[dst][node] = reach[dst][node] || reach[src][node]
			}
			indegree[dst]--
			if indegree[dst] == 0 {
				queue.PushBack(dst)
			}
		}
	}
	answer := make([]bool, len(queries))
	var u, v int
	for idx := range queries {
		u, v = queries[idx][0], queries[idx][1]
		answer[idx] = reach[v][u]
	}
	return answer
}
