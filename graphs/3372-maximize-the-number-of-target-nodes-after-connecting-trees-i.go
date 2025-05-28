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
		var cnt = 1
		for _, nei := range adj[node] {
			if nei != parent {
				cnt += dfs(node, nei, adj, k-1)
			}
		}
		return cnt
	}
	var traverse = func(edges [][]int, k int, val int) ([]int, int) {
		var n = len(edges) + 1
		var adj = make([][]int, n)
		var a, b int
		for idx := range edges {
			a, b = edges[idx][0], edges[idx][1]
			adj[a] = append(adj[a], b)
			adj[b] = append(adj[b], a)
		}
		var answer = make([]int, n)
		var mx = 0
		for node := 0; node < n; node++ {
			answer[node] = dfs(-1, node, adj, k) + val
			mx = max(mx, answer[node])
		}
		return answer, mx
	}
	var _, max2 = traverse(edges2, k-1, 0)
	var answer, _ = traverse(edges1, k, max2)
	return answer
}
