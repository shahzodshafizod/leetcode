package dp

// https://leetcode.com/problems/solving-questions-with-brainpower/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(n)
// Space: O(n)
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n)
	dp[n-1] = int64(questions[n-1][0])
	var nextIdx int
	for idx := n - 2; idx >= 0; idx-- {
		dp[idx] = int64(questions[idx][0])
		nextIdx = idx + questions[idx][1] + 1
		if nextIdx < n {
			dp[idx] += dp[nextIdx]
		}
		dp[idx] = max(dp[idx], dp[idx+1])
	}
	return dp[0]
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(n)
// // Space: O(n)
// func mostPoints(questions [][]int) int64 {
// 	var n = len(questions)
// 	var memo = make([]*int64, n)
// 	var dp func(idx int) int64
// 	dp = func(idx int) int64 {
// 		if idx >= n {
// 			return 0
// 		}
// 		if memo[idx] != nil {
// 			return *memo[idx]
// 		}
// 		memo[idx] = new(int64)
// 		*memo[idx] = max(
// 			dp(idx+1),
// 			int64(questions[idx][0])+dp(idx+questions[idx][1]+1),
// 		)
// 		return *memo[idx]
// 	}
// 	return dp(0)
// }
