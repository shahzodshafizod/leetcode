package graphs

import "slices"

// https://leetcode.com/problems/parallel-courses-iii/

// Approach: Breadth-First Search
// Time: O(E+V), E=# of edges, V=# of nodes/vertices
// Space: O(E+V)
func minimumTime(n int, relations [][]int, time []int) int {
	var adjList = make([][]int, n)
	var indegree = make([]int, n)
	var src, dst int
	for idx := range relations {
		src, dst = relations[idx][0]-1, relations[idx][1]-1
		adjList[src] = append(adjList[src], dst)
		indegree[dst]++
	}
	var queue = make([]int, 0)
	for node, degree := range indegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}
	var dist = make([]int, n)
	copy(dist, time)
	var idx, size = 0, len(queue)
	for idx < size {
		src = queue[idx]
		idx++
		for _, dst := range adjList[src] {
			dist[dst] = max(dist[dst], dist[src]+time[dst])
			indegree[dst]--
			if indegree[dst] == 0 {
				queue = append(queue, dst)
				size++
			}
		}
	}
	return slices.Max(dist)
}

// // Approach: Depth-First Search
// // Time: O(E+V), E=# of edges, V=# of nodes/vertices
// // Space: O(E+V)
// func minimumTime(n int, relations [][]int, time []int) int {
// 	var adjList = make([][]int, n)
// 	var src, dst int
// 	for idx := range relations {
// 		src, dst = relations[idx][0]-1, relations[idx][1]-1
// 		adjList[src] = append(adjList[src], dst)
// 	}
// 	var memo = make([]*int, n)
// 	var dfs func(src int) int
// 	dfs = func(src int) int {
// 		if memo[src] != nil {
// 			return *memo[src]
// 		}
// 		var months = time[src]
// 		for _, dst := range adjList[src] {
// 			months = max(months, time[src]+dfs(dst))
// 		}
// 		memo[src] = &months
// 		return months
// 	}
// 	var months = 0
// 	for node := 0; node < n; node++ {
// 		months = max(months, dfs(node))
// 	}
// 	return months
// }
