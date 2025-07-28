package graphs

// https://leetcode.com/problems/jump-game-v/

// Top-Down Dynamic Programming
// Time: O(n x d)
// Space: O(n)
func maxJumps(arr []int, d int) int {
	n := len(arr)
	dp := make([]int, n)

	var dfs func(idx int) int

	dfs = func(idx int) int {
		if dp[idx] != 0 {
			return dp[idx]
		}

		dp[idx] = 1
		// move left
		for step := 1; step <= d && idx-step >= 0 && arr[idx] > arr[idx-step]; step++ {
			dp[idx] = max(dp[idx], 1+dfs(idx-step))
		}
		// move right
		for step := 1; step <= d && idx+step < n && arr[idx] > arr[idx+step]; step++ {
			dp[idx] = max(dp[idx], 1+dfs(idx+step))
		}

		return dp[idx]
	}
	maxJump := 0

	for idx := 0; idx < n; idx++ {
		maxJump = max(maxJump, dfs(idx))
	}

	return maxJump
}

// // Approach: Monotonic Stack + DFS on Graph
// // Time: O(3n) = O(n)
// // Space: O(3n) = O(n) (adjList, stack, dp)
// func maxJumps(arr []int, d int) int {
// 	var n = len(arr)
// 	var adjList = make([][]int, n)
// 	var makeGraph = func(from int, to int, step int) {
// 		var stack = make([]int, n)
// 		var slen = 0
// 		var neighbor int
// 		for curr := from; curr != to; curr += step {
// 			for slen > 0 && arr[curr] > arr[stack[slen-1]] {
// 				slen--
// 				neighbor = stack[slen]
// 				if step*(curr-neighbor) <= d {
// 					adjList[curr] = append(adjList[curr], neighbor)
// 				}
// 			}
// 			stack[slen] = curr
// 			slen++
// 		}
// 	}
// 	makeGraph(0, n, 1)     // add prev neighbors
// 	makeGraph(n-1, -1, -1) // add next neighbors
// 	var dp = make([]int, n)
// 	var dfs func(idx int) int
// 	dfs = func(idx int) int {
// 		if dp[idx] != 0 {
// 			return dp[idx]
// 		}
// 		dp[idx] = 1
// 		for _, next := range adjList[idx] {
// 			dp[idx] = max(dp[idx], 1+dfs(next))
// 		}
// 		return dp[idx]
// 	}
// 	var maxJump = 0
// 	for idx := 0; idx < n; idx++ {
// 		maxJump = max(maxJump, dfs(idx))
// 	}
// 	return maxJump
// }
