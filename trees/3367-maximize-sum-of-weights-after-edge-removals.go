package trees

import (
	"slices"
)

// https://leetcode.com/problems/maximize-sum-of-weights-after-edge-removals/

// Approach #2: Depth-First Search
// Time: O(n)
// Space: O(n)
func maximizeSumOfWeights(edges [][]int, k int) int64 {
	n := len(edges) + 1
	graph := make([][][2]int, n)

	var (
		weights int64
		u, v, w int
	)

	valid := true

	for _, edge := range edges {
		u, v, w = edge[0], edge[1], edge[2]

		graph[u] = append(graph[u], [2]int{v, w})
		graph[v] = append(graph[v], [2]int{u, w})

		if valid && (len(graph[u]) > k || len(graph[v]) > k) {
			valid = false
		}

		weights += int64(w)
	}

	if valid {
		return weights
	}

	var dfs func(parent int, node int) (int64, int64)

	dfs = func(parent, node int) (int64, int64) {
		var (
			res, skip, choose int64
			nnode, weight     int
			gains             []int64
		)

		for _, nei := range graph[node] {
			nnode, weight = nei[0], nei[1]
			if nnode == parent {
				continue
			}

			skip, choose = dfs(node, nnode)
			res += choose
			gains = append(gains, max(0, skip+int64(weight)-choose))
		}

		// sort in descending order
		slices.SortFunc(gains, func(a int64, b int64) int { return int(b - a) })

		for i := range min(len(gains), k-1) {
			res += gains[i]
		}

		choose = res
		skip = res

		if len(gains) >= k {
			choose += gains[k-1]
		}

		return skip, choose
	}

	_, choose := dfs(-1, 0)

	return choose
}

// // Approach #1: Brute-Force (Time Limit Exceeded)
// // Time: O(nn)
// // Space: O(n)
// func maximizeSumOfWeights(edges [][]int, k int) int64 {
// 	n := len(edges) + 1
// 	degrees := make([]int, n) // for nodes

// 	var u, v int
// 	for _, edge := range edges {
// 		u, v = edge[0], edge[1]
// 		degrees[u]++
// 		degrees[v]++
// 	}

// 	deleted := make([]bool, n-1) // for edges

// 	var dp func(eid int) int64

// 	dp = func(eid int) int64 {
// 		var (
// 			weights int64
// 			u, v, w int
// 		)

// 		for idx, edge := range edges {
// 			u, v, w = edge[0], edge[1], edge[2]

// 			if deleted[idx] {
// 				continue
// 			}

// 			if degrees[u] > k || degrees[v] > k {
// 				weights = 0

// 				break
// 			}

// 			weights += int64(w)
// 		}

// 		if eid == n-1 {
// 			return weights
// 		}

// 		u, v = edges[eid][0], edges[eid][1]

// 		// 1. skip
// 		weights = max(weights, dp(eid+1))

// 		// 2. delete
// 		deleted[eid] = true
// 		degrees[u]--
// 		degrees[v]--
// 		weights = max(weights, dp(eid+1))
// 		deleted[eid] = false
// 		degrees[u]++
// 		degrees[v]++

// 		return weights
// 	}

// 	return dp(0)
// }
