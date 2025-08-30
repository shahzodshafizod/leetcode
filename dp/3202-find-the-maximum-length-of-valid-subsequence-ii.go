package dp

// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/

func maximumLength3202(nums []int, k int) int {
	dp := make([][]int, k)
	for idx := range dp {
		dp[idx] = make([]int, k)
	}

	length := 0

	for _, curr := range nums {
		curr %= k
		for prev := range k {
			dp[curr][prev] = dp[prev][curr] + 1
			length = max(length, dp[curr][prev])
		}
	}

	return length
}
