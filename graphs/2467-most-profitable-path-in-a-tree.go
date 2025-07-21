package graphs

import "math"

// https://leetcode.com/problems/most-profitable-path-in-a-tree/

// Approach: Depth-First Search
// Time: O(n)
// Space: O(n)
func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges) + 1
	graph := make([][]int, n)
	var a, b int
	for idx := range edges {
		a, b = edges[idx][0], edges[idx][1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	var dfs func(parent int, node int, adist int) (int, int)
	// returns (alice profit, bob dist)
	dfs = func(parent int, node int, adist int) (int, int) {
		profit := math.MinInt
		bdist := -1
		for _, nei := range graph[node] {
			if nei != parent {
				p, d := dfs(node, nei, adist+1)
				profit = max(profit, p+amount[node])
				bdist = max(bdist, d)
			}
		}
		if profit == math.MinInt {
			profit = amount[node]
		}
		if node == bob {
			bdist = 0
		}
		if bdist != -1 {
			if bdist == adist {
				profit -= amount[node] / 2
			} else if bdist < adist {
				profit -= amount[node]
			}
			bdist++
		}
		return profit, bdist
	}
	profit, _ := dfs(-1, 0, 0)
	return profit
}
