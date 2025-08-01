package graphs

// https://leetcode.com/problems/maximum-path-quality-of-a-graph/

// Approach: Backtracking
// Time: O(n∗(4^n))
// Space: O(n)
func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	n := len(values)
	adj := make([][][2]int, n)

	var u, v, t int
	for _, edge := range edges {
		u, v, t = edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], [2]int{v, t})
		adj[v] = append(adj[v], [2]int{u, t})
	}

	var dfs func(node int, value int, time int) int

	dfs = func(node int, value int, time int) int {
		var best int
		if node == 0 {
			best = value
		}

		oldValue := values[node]

		values[node] = 0
		for _, nei := range adj[node] {
			if time >= nei[1] {
				best = max(best, dfs(nei[0], value+values[nei[0]], time-nei[1]))
			}
		}

		values[node] = oldValue

		return best
	}

	return dfs(0, values[0], maxTime)
}
