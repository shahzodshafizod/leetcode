package graphs

// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/

// Approach: Breadth-First Search
// Time: O(26n) = O(n)
// Space: O(26n) = O(n)
func largestPathValue(colors string, edges [][]int) int {
	var n = len(colors)
	var adj = make([][]int, n)
	var indegree = make([]int, n)
	var src, dst int
	for idx := range edges {
		src, dst = edges[idx][0], edges[idx][1]
		adj[src] = append(adj[src], dst)
		indegree[dst] += 1
	}
	var queue = make([]int, 0)
	var colids = make([]int, n)
	for node := 0; node < n; node++ {
		if indegree[node] == 0 {
			queue = append(queue, node)
		}
		colids[node] = int(colors[node] - 'a')
	}
	var dp = make([][26]int, n)
	var result = 0
	for size := len(queue); size > 0; size = len(queue) {
		for idx := 0; idx < size; idx++ {
			var curr = queue[idx]
			n--
			dp[curr][colids[curr]]++
			for c := 0; c < 26; c++ {
				result = max(result, dp[curr][c])
			}
			for _, next := range adj[curr] {
				for c := 0; c < 26; c++ {
					dp[next][c] = max(dp[next][c], dp[curr][c])
				}
				indegree[next]--
				if indegree[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
		queue = queue[size:]
	}
	if n != 0 {
		return -1
	}
	return result
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
// 	var dp = make([][]int, n)
// 	for idx := 0; idx < n; idx++ {
// 		colids[idx] = int(colors[idx] - 'a')
// 		dp[idx] = make([]int, 26)
// 	}
// 	var extra int
// 	var memo = make([]*int, n)
// 	var dfs func(node int) int
// 	dfs = func(node int) int {
// 		if memo[node] != nil && *memo[node] == -1 {
// 			return math.MaxInt
// 		}
// 		if memo[node] != nil {
// 			return *memo[node]
// 		}
// 		memo[node] = new(int)
// 		*memo[node] = -1 // sign of detecting cycles
// 		dp[node][colids[node]] = 1
// 		for _, next := range adj[node] {
// 			if dfs(next) == math.MaxInt {
// 				return math.MaxInt
// 			}
// 			for c := 0; c < 26; c++ {
// 				extra = 0
// 				if c == colids[node] {
// 					extra = 1
// 				}
// 				dp[node][c] = max(
// 					dp[node][c],
// 					dp[next][c]+extra,
// 				)
// 			}
// 		}
// 		*memo[node] = slices.Max(dp[node])
// 		return *memo[node]
// 	}
// 	var count = 0
// 	for node := 0; node < n; node++ {
// 		count = max(count, dfs(node))
// 	}
// 	if count == math.MaxInt {
// 		return -1
// 	}
// 	return count
// }
