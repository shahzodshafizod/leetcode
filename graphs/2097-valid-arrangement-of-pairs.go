package graphs

// https://leetcode.com/problems/valid-arrangement-of-pairs/

/*
Eulerian paths have a couple of conditions:
	- In an undirected graph, either all nodes have an even degree,
		or exactly two have an odd degree.
	- In a directed graph (which is what we have here), we need to check if:
		- Each node’s degree matches its degree.
		- Or, exactly one node has one more outgoing edge
			(degree = degree + 1), which indicates our starting point.
https://www.youtube.com/watch?v=8MpoO2zA2l4
*/

// Approach 2: Hierholzer's Algorithm (Iterative)
// Time: O(V+E), V=# of vertices(nodes), E=# of edges
// Space: O(V+E)
func validArrangement(pairs [][]int) [][]int {
	adjList := make(map[int][]int)
	degree := make(map[int]int)

	var start, end int
	for idx := range pairs {
		start, end = pairs[idx][0], pairs[idx][1]
		adjList[start] = append(adjList[start], end)
		degree[start]++
		degree[end]--
	}

	start = pairs[0][0]

	for node := range degree {
		if degree[node] == 1 {
			start = node
			break
		}
	}

	stack := []int{start}

	var curr, next int

	order := make([]int, 0, len(pairs)+1)

	for size := len(stack); size > 0; size = len(stack) {
		curr = stack[size-1]
		if l := len(adjList[curr]); l > 0 {
			next = adjList[curr][l-1]
			adjList[curr] = adjList[curr][:l-1]

			stack = append(stack, next)
		} else {
			stack = stack[:size-1]

			order = append(order, curr)
		}
	}

	arrangement := make([][]int, 0, len(pairs))
	for i := len(order) - 1; i > 0; i-- {
		arrangement = append(arrangement, []int{order[i], order[i-1]})
	}

	return arrangement
}

// // Approach 1: Eulerian Path (Recursive)
// // Time: O(V+E), V=# of vertices(nodes), E=# of edges
// // Space: O(V+E)
// func validArrangement(pairs [][]int) [][]int {
// 	var adjList = make(map[int][]int)
// 	var degree = make(map[int]int)
// 	var start, end int
// 	for idx := range pairs {
// 		start, end = pairs[idx][0], pairs[idx][1]
// 		adjList[start] = append(adjList[start], end)
// 		degree[start]++
// 		degree[end]--
// 	}
// 	start = pairs[0][0]
// 	for node := range degree {
// 		if degree[node] == 1 {
// 			start = node
// 			break
// 		}
// 	}
// 	var order = make([]int, 0, len(pairs)+1)
// 	var visit func(node int)
// 	visit = func(node int) {
// 		for size := len(adjList[node]); size > 0; size = len(adjList[node]) {
// 			next := adjList[node][size-1]
// 			adjList[node] = adjList[node][:size-1]
// 			visit(next)
// 		}
// 		order = append(order, node)
// 	}
// 	visit(start)
// 	var arrangement = make([][]int, 0, len(pairs))
// 	for idx := len(order) - 2; idx >= 0; idx-- {
// 		arrangement = append(arrangement, []int{order[idx+1], order[idx]})
// 	}
// 	return arrangement
// }
