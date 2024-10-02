package dp

// https://leetcode.com/problems/burst-balloons/

// Approach: Top-Down Dynamic Programming
// Time: O(N^3)
// Space: O(N^2)
func maxCoins(nums []int) int {
	var n = len(nums) + 2
	var bordered = make([]int, n)
	copy(bordered[1:], nums)
	nums = bordered
	nums[0], nums[n-1] = 1, 1
	var dp = make([][]int, n)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}
	for left := n - 2; left > 0; left-- {
		for right := left; right < n-1; right++ {
			for idx := left; idx <= right; idx++ {
				dp[left][right] = max(
					dp[left][right],
					nums[left-1]*nums[idx]*nums[right+1]+
						dp[left][idx-1]+
						dp[idx+1][right],
				)
			}
		}
	}
	// [0] and [n-1] are not part of the nums
	return dp[1][n-2]
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(N^3)
// // Space: O(N^2)
// func maxCoins(nums []int) int {
// 	var n = len(nums)
// 	var dp = make([][]*int, n)
// 	for idx := range dp {
// 		dp[idx] = make([]*int, n)
// 	}
// 	var dfs func(prev int, left int, right int, next int) int
// 	dfs = func(prev int, left int, right int, next int) int {
// 		if left > right {
// 			return 0
// 		}
// 		if dp[left][right] != nil {
// 			return *dp[left][right]
// 		}
// 		dp[left][right] = new(int)
// 		var coins int
// 		for idx := left; idx <= right; idx++ {
// 			coins = prev*nums[idx]*next +
// 				dfs(prev, left, idx-1, nums[idx]) +
// 				dfs(nums[idx], idx+1, right, next)
// 			*dp[left][right] = max(*dp[left][right], coins)
// 		}
// 		return *dp[left][right]
// 	}
// 	return dfs(1, 0, n-1, 1)
// }
