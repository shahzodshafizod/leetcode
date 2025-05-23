package heaps

import (
	"sort"
)

// https://leetcode.com/problems/find-number-of-coins-to-place-in-tree-nodes/

// Approach: Depth-First Search + Sorting
// Time: O(n)
// Space: O(n)
func placedCoins(edges [][]int, cost []int) []int64 {
	var n = len(cost)
	var graph = make([][]int, n)
	var a, b int
	for _, edge := range edges {
		a, b = edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	var coin = make([]int64, n)
	var dfs func(parent int, node int) []int
	dfs = func(parent int, node int) []int {
		var cs = []int{cost[node]}
		for _, child := range graph[node] {
			if child != parent {
				cs = append(cs, dfs(node, child)...)
			}
		}
		sort.Ints(cs) // O(5 log 5)
		var n = len(cs)
		if n > 5 {
			// [neg, neg, pos, pos, pos]
			cs = []int{cs[0], cs[1], cs[n-3], cs[n-2], cs[n-1]}
			n = 5
		}
		coin[node] = 1
		if n >= 3 {
			coin[node] = int64(max(0, cs[0]*cs[1]*cs[n-1], cs[n-3]*cs[n-2]*cs[n-1]))
		}
		return cs
	}
	dfs(-1, 0)
	return coin
}
