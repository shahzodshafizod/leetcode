package unionfinds

// https://leetcode.com/problems/redundant-connection/

// Approach: Union Find
// Time: O(N⋅α(N))
// Space: O(n)
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	for node := 0; node <= n; node++ {
		parent[node] = node
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x int, y int) bool {
		px := find(x)
		py := find(y)
		if px == py {
			return false
		}
		parent[py] = px
		return true
	}
	var edge []int
	for idx := range edges {
		if !union(edges[idx][0], edges[idx][1]) {
			edge = edges[idx]
		}
	}
	return edge
}
