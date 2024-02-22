package dp

/*
  |a b c d e|
--+-+-+-+-+-+
a |1|1|1|1|1|
c |1|1|2|2|2|
e |1|1|2|2|3|
*/

// https://leetcode.com/problems/longest-common-subsequence/

// bottom up
func longestCommonSubsequence(text1 string, text2 string) int {
	var len1, len2 = len(text1), len(text2)
	var dp = make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len1][len2]
}

// // Time Limit Exceeded
// // bottom up
// func longestCommonSubsequence(text1 string, text2 string) int {
// 	var len1, len2 = len(text1), len(text2)
// 	var memo = make([][]int, len1)
// 	for i := 0; i < len1; i++ {
// 		memo[i] = make([]int, len2)
// 	}
// 	var dp func(row int, col int) int
// 	dp = func(row int, col int) int {
// 		if row == len1 || col == len2 {
// 			return 0
// 		}
// 		if memo[row][col] != 0 {
// 			return memo[row][col]
// 		}
// 		if text1[row] == text2[col] {
// 			memo[row][col] = 1 + dp(row+1, col+1)
// 		} else {
// 			memo[row][col] = max(dp(row+1, col), dp(row, col+1))
// 		}
// 		return memo[row][col]
// 	}
// 	return dp(0, 0)
// }
