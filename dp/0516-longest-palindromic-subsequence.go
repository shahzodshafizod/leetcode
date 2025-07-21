package dp

// https://leetcode.com/problems/longest-palindromic-subsequence/

// time: O(2^n) w/o memoization
// time: O(n^2) w/ memoization
func longestPalindromeSubseq(s string) int {
	len := len(s)
	memo := make([][]int, len)
	for i := 0; i < len; i++ {
		memo[i] = make([]int, len)
	}
	var dp func(left int, right int) int
	dp = func(left int, right int) int {
		if left > right {
			return 0
		}
		if left == right {
			return 1
		}
		if memo[left][right] != 0 {
			return memo[left][right]
		}
		if s[left] == s[right] {
			return 2 + dp(left+1, right-1)
		}
		memo[left+1][right] = dp(left+1, right)
		memo[left][right-1] = dp(left, right-1)
		memo[left][right] = max(memo[left+1][right], memo[left][right-1])
		return memo[left][right]
	}
	return dp(0, len-1)
}

/*
another solution is using 1143-longest-common-subsequence:
longestCommonSubsequence(s, s.reverse())
*/
