package graphs

// https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-ii/

// Approach: Depth-First Search
// Time: O(n+m)
// Space: O(n+m)
func maxTargetNodes3373(edges1 [][]int, edges2 [][]int) []int {
	var dfs func(parent int, node int, adj [][]int, color []int, col int) int
	dfs = func(parent int, node int, adj [][]int, color []int, col int) int {
		color[node] = col
		targets := col
		for _, nei := range adj[node] {
			if nei != parent {
				targets += dfs(node, nei, adj, color, col^1)
			}
		}
		return targets
	}

	makeAdjList := func(edges [][]int, n int) [][]int {
		adj := make([][]int, n)
		var a, b int
		for idx := range edges {
			a, b = edges[idx][0], edges[idx][1]
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}
		return adj
	}

	m := len(edges2) + 1
	evens2 := dfs(-1, 0, makeAdjList(edges2, m), make([]int, m), 1)
	max2 := max(evens2, m-evens2)

	n := len(edges1) + 1
	color1 := make([]int, n)
	evens1 := dfs(-1, 0, makeAdjList(edges1, n), color1, 1)
	count1 := [2]int{n - evens1, evens1}

	answer := make([]int, n)
	for node := 0; node < n; node++ {
		answer[node] = count1[color1[node]] + max2
	}
	return answer
}
