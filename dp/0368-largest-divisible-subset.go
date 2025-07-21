package dp

import "sort"

// https://leetcode.com/problems/largest-divisible-subset/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(nn)
// Space: O(nn)
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([][]int, n)
	var answer []int
	for idx := n - 1; idx >= 0; idx-- {
		dp[idx] = append(dp[idx], nums[idx])
		for j := idx + 1; j < n; j++ {
			if nums[j]%nums[idx] == 0 && len(dp[j])+1 > len(dp[idx]) {
				dp[idx] = append([]int{nums[idx]}, dp[j]...)
			}
		}
		if len(dp[idx]) > len(answer) {
			answer = dp[idx]
		}
	}
	return answer
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(nn)
// // Space: O(nn)
// func largestDivisibleSubset(nums []int) []int {
// 	sort.Ints(nums)
// 	var n = len(nums)
// 	var memo = make([]*[]int, n)
// 	var dp func(idx int) []int
// 	dp = func(idx int) []int {
// 		if idx == n {
// 			return nil
// 		}
// 		if memo[idx] != nil {
// 			return *memo[idx]
// 		}
// 		var answer []int
// 		for j := idx + 1; j < n; j++ {
// 			if nums[j]%nums[idx] == 0 {
// 				var tmp = dp(j)
// 				if len(tmp) > len(answer) {
// 					answer = tmp
// 				}
// 			}
// 		}
// 		memo[idx] = new([]int)
// 		*memo[idx] = append([]int{nums[idx]}, answer...)
// 		return *memo[idx]
// 	}
// 	var answer []int
// 	for idx := 0; idx < n; idx++ {
// 		if len(dp(idx)) > len(answer) {
// 			answer = *memo[idx]
// 		}
// 	}
// 	return answer
// }
