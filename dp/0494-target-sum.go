package dp

// https://leetcode.com/problems/target-sum/

// Approach #4: Bottom-Up Dynamic Programming (Space Optimized)
// Time: O(n*sum(nums))
// Space: O(sum(nums))
func findTargetSumWays(nums []int, target int) int {
	var n = len(nums)
	var dp = make(map[int]int)
	dp[0] = 1
	for idx := 0; idx < n; idx++ {
		prev := dp
		dp = make(map[int]int)
		for currSum, count := range prev {
			dp[currSum+nums[idx]] += count
			dp[currSum-nums[idx]] += count
		}
	}
	return dp[target]
}

// // Approach #3: Bottom-Up Dynamic Programming
// // Time: O(n*sum(nums))
// // Space: O(n*sum(nums))
// func findTargetSumWays(nums []int, target int) int {
// 	var n = len(nums)
// 	var dp = make([]map[int]int, n+1)
// 	dp[0] = make(map[int]int)
// 	dp[0][0] = 1
// 	for idx := 0; idx < n; idx++ {
// 		dp[idx+1] = make(map[int]int)
// 		for currSum, count := range dp[idx] {
// 			dp[idx+1][currSum+nums[idx]] += count
// 			dp[idx+1][currSum-nums[idx]] += count
// 		}
// 	}
// 	return dp[n][target]
// }

// // Approach #2: Backtracking with Memoization
// // Time: O(n*sum(nums))
// // Space: O(n*sum(nums))
// func findTargetSumWays(nums []int, target int) int {
// 	var n = len(nums)
// 	var memo = make([]map[int]int, n)
// 	var backtrack func(idx int, result int) int
// 	backtrack = func(idx int, result int) int {
// 		if idx == n {
// 			if result == target {
// 				return 1
// 			}
// 			return 0
// 		}
// 		if memo[idx] == nil {
// 			memo[idx] = make(map[int]int)
// 		}
// 		if ways, found := memo[idx][result]; found {
// 			return ways
// 		}
// 		memo[idx][result] = backtrack(idx+1, result+nums[idx]) +
// 			backtrack(idx+1, result-nums[idx])
// 		return memo[idx][result]
// 	}
// 	return backtrack(0, 0)
// }

// // Approach #1: Backtracking
// // Time: O(2^n)
// // Space: O(n), for recursion stack
// func findTargetSumWays(nums []int, target int) int {
// 	var n = len(nums)
// 	var backtrack func(idx int, result int) int
// 	backtrack = func(idx int, result int) int {
// 		if idx == n {
// 			if result == target {
// 				return 1
// 			}
// 			return 0
// 		}
// 		return backtrack(idx+1, result+nums[idx]) +
// 			backtrack(idx+1, result-nums[idx])
// 	}
// 	return backtrack(0, 0)
// }
