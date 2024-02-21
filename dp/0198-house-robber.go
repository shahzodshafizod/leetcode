package dp

// https://leetcode.com/problems/house-robber/

// Iterative (bottom-up)
func rob(nums []int) int {
	var n = len(nums)
	var prev = 0
	for i := 2; i < n; i++ {
		nums[i] += max(nums[i-2], prev)
		prev = nums[i-2]
	}
	if n > 1 {
		prev = nums[n-2]
	}
	return max(prev, nums[n-1])
}

// // Time Limit Exceeded
// // Recursive + memo (memoization) (top-down)
// func rob(nums []int) int {
// 	var n = len(nums)
// 	var dp func(idx int, memo []int) int
// 	dp = func(idx int, memo []int) int {
// 		if idx >= n {
// 			return 0
// 		}
// 		if idx >= n-2 {
// 			return nums[idx]
// 		}
// 		if memo[idx] > 0 {
// 			return memo[idx]
// 		}
// 		memo[idx] = nums[idx] + max(dp(idx+2, memo), dp(idx+3, memo))
// 		return memo[idx]
// 	}
// 	var memo = make([]int, n)
// 	return max(dp(0, memo), dp(1, memo))
// }
