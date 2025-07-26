package graphs

// https://leetcode.com/problems/minimum-score-after-removals-on-a-tree/

func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	adj := make([][]int, n)
	var u, v int
	for _, edge := range edges {
		u, v = edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	total := 0
	for _, num := range nums {
		total ^= num
	}
	tin := make([]int, n)  // time in
	tout := make([]int, n) // time out
	nodeXor := make([]int, n)
	var dfs func(parent int, node int, timer int) int
	dfs = func(parent int, node int, timer int) int {
		tin[node] = timer
		timer++
		nodeXor[node] = nums[node]
		for _, nxt := range adj[node] {
			if nxt != parent {
				timer = dfs(node, nxt, timer)
				nodeXor[node] ^= nodeXor[nxt]
			}
		}
		tout[node] = timer
		return timer + 1
	}
	dfs(-1, 0, 0)
	isAncestor := func(node1 int, node2 int) bool {
		return tin[node1] < tin[node2] && tout[node1] > tout[node2]
	}
	res := (1 << 32) - 1
	var a, b, c, d, x1, x2, x3 int
	for i, m := 0, len(edges); i < m-1; i++ {
		a, b = edges[i][0], edges[i][1]
		if isAncestor(a, b) {
			a, b = b, a
		}
		for j := i + 1; j < m; j++ {
			c, d = edges[j][0], edges[j][1]
			if isAncestor(c, d) {
				c, d = d, c
			}

			x1, x2, x3 = 0, 0, 0
			if isAncestor(a, c) {
				x1 = total ^ nodeXor[a]
				x2 = nodeXor[a] ^ nodeXor[c]
				x3 = nodeXor[c]
			} else if isAncestor(c, a) {
				x1 = total ^ nodeXor[c]
				x2 = nodeXor[c] ^ nodeXor[a]
				x3 = nodeXor[a]
			} else {
				x1 = total ^ nodeXor[a] ^ nodeXor[c]
				x2 = nodeXor[a]
				x3 = nodeXor[c]
			}

			res = min(res, max(x1, x2, x3)-min(x1, x2, x3))
		}
	}
	return res
}
