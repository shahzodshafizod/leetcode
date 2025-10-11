package graphs

// https://leetcode.com/problems/longest-cycle-in-a-graph/

// Hint: Each node can be part of at most one cycle

func longestCycle(edges []int) int {
	n := len(edges)
	dist := make([]int, n)

	var dfs func(node int, length int) int

	dfs = func(node int, length int) int {
		if node == -1 || dist[node] == -1 {
			return -1
		}

		if 0 < dist[node] && dist[node] < length {
			return length - dist[node]
		}

		dist[node] = length
		length = dfs(edges[node], length+1)
		dist[node] = -1

		return length
	}

	length := -1

	for node := range n {
		if dist[node] == 0 {
			length = max(length, dfs(node, 1))
		}
	}

	return length
}
