package dp

// https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/

// Approach #2: Bottom-Up Dynamic Programming: Longest Common Sequence
// Time: O(nn)
// Space: O(nn)
func minInsertions(s string) int {
	n := len(s)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range n {
		for j := range n {
			if s[i] == s[n-j-1] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return n - dp[n][n]
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(nn)
// // Space: O(nn)
// func minInsertions(s string) int {
// 	n := len(s)

// 	memo := make([][]*int, n)
// 	for i := range memo {
// 		memo[i] = make([]*int, n)
// 	}

// 	var dp func(left int, right int) int

// 	dp = func(left, right int) int {
// 		if left >= right {
// 			return 0
// 		}

// 		if memo[left][right] != nil {
// 			return *memo[left][right]
// 		}

// 		memo[left][right] = new(int)
// 		if s[left] == s[right] {
// 			*memo[left][right] = dp(left+1, right-1)
// 		} else {
// 			*memo[left][right] = 1 + min(dp(left+1, right), dp(left, right-1))
// 		}

// 		return *memo[left][right]
// 	}

// 	return dp(0, n-1)
// }
