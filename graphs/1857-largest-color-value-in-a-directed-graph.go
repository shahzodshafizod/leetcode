package graphs

// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/

// Approach: Breadth-First Search
// Time: O(26n) = O(n)
// Space: O(26n) = O(n)
func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	adj := make([][]int, n)
	indegrees := make([]int, n)

	var src, dst int
	for idx := range edges {
		src, dst = edges[idx][0], edges[idx][1]
		adj[src] = append(adj[src], dst)
		indegrees[dst] += 1
	}

	queue := make([]int, n)
	size := 0
	colids := make([]int, n)

	for node := 0; node < n; node++ {
		if indegrees[node] == 0 {
			queue[size] = node
			size++
		}

		colids[node] = int(colors[node] - 'a')
	}

	dp := make([][26]int, n)
	value := 0

	var node int
	for size > 0 {
		node = queue[size-1]
		size--
		n--

		dp[node][colids[node]]++
		for _, cnt := range dp[node] {
			value = max(value, cnt)
		}

		for _, next := range adj[node] {
			for c := 0; c < 26; c++ {
				dp[next][c] = max(dp[next][c], dp[node][c])
			}

			indegrees[next]--
			if indegrees[next] == 0 {
				queue[size] = next
				size++
			}
		}
	}

	if n != 0 {
		return -1
	}

	return value
}

// // Approach: Depth-First Search
// // Time: O(26n) = O(n)
// // Space: O(26n) = O(n)
// func largestPathValue(colors string, edges [][]int) int {
// 	var n = len(colors)
// 	var adj = make([][]int, n)
// 	var src, dst int
// 	for idx := range edges {
// 		src, dst = edges[idx][0], edges[idx][1]
// 		adj[src] = append(adj[src], dst)
// 	}
// 	var colids = make([]int, n)
// 	for idx := 0; idx < n; idx++ {
// 		colids[idx] = int(colors[idx] - 'a')
// 	}
// 	var count = make([][26]int, n)
// 	var memo = make([]*int, n)
// 	var dfs func(node int) int
// 	dfs = func(node int) int {
// 		if memo[node] != nil {
// 			return *memo[node]
// 		}
// 		var color = colids[node]
// 		memo[node] = new(int)
// 		*memo[node] = n + 1
// 		var value = 1
// 		for _, next := range adj[node] {
// 			if dfs(next) > n {
// 				return n + 1
// 			}
// 			for c := 0; c < 26; c++ {
// 				count[node][c] = max(
// 					count[node][c],
// 					count[next][c],
// 				)
// 				value = max(value, count[node][c])
// 			}
// 		}
// 		count[node][color]++
// 		value = max(value, count[node][color])
// 		memo[node] = &value
// 		return value
// 	}
// 	var value = 0
// 	for node := 0; node < n; node++ {
// 		value = max(value, dfs(node))
// 	}
// 	if value > n {
// 		return -1
// 	}
// 	return value
// }
