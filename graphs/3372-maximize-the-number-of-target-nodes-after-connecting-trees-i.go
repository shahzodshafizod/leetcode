package graphs

// https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/

// Approach: Depth-First Search
// Time: O(nn+mm)
// Space: O(n+m)
func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	var dfs func(parent int, node int, adj [][]int, k int) int

	dfs = func(parent int, node int, adj [][]int, k int) int {
		if k < 0 {
			return 0
		}

		cnt := 1

		for _, nei := range adj[node] {
			if nei != parent {
				cnt += dfs(node, nei, adj, k-1)
			}
		}

		return cnt
	}
	traverse := func(edges [][]int, k int, val int) ([]int, int) {
		n := len(edges) + 1
		adj := make([][]int, n)

		var a, b int
		for idx := range edges {
			a, b = edges[idx][0], edges[idx][1]
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}

		answer := make([]int, n)
		mx := 0

		for node := range n {
			answer[node] = dfs(-1, node, adj, k) + val
			mx = max(mx, answer[node])
		}

		return answer, mx
	}
	_, max2 := traverse(edges2, k-1, 0)
	answer, _ := traverse(edges1, k, max2)

	return answer
}
