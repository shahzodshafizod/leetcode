package dp

import "sort"

// https://leetcode.com/problems/maximize-consecutive-elements-in-an-array-after-modification/

// Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(n log n)
// Space: O(1)
func maxSelectedElements(nums []int) int {
	sort.Ints(nums)

	dp := make(map[int]int)

	var res int
	for _, num := range nums {
		newdp := make(map[int]int)
		// keep the number the same
		newdp[num] = dp[num-1] + 1
		// increase the number
		newdp[num+1] = dp[num] + 1
		// copy the old count
		newdp[num-1] = dp[num-1]
		dp = newdp
		res = max(res, dp[num], dp[num+1])
	}

	return res
}

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(n log n)
// // Space: O(n)
// func maxSelectedElements(nums []int) int {
// 	sort.Ints(nums)
// 	dp := make(map[int]int, len(nums))
// 	var res int
// 	for _, num := range nums {
// 		dp[num+1] = dp[num] + 1
// 		dp[num] = dp[num-1] + 1
// 		res = max(res, dp[num+1], dp[num])
// 	}
// 	return res
// }

// // Approach #1: Top-Down Dynamic Programming (Memory Limit Exceeded)
// // Time: O(nm), n=len(nums), m=max(nums)
// // Space: O(mn)
// func maxSelectedElements(nums []int) int {
// 	sort.Ints(nums)
// 	n := len(nums)
// 	memo := make([]map[int]*int, n)
// 	for idx := range memo {
// 		memo[idx] = make(map[int]*int)
// 	}
// 	var dp func(idx int, prev int) int
// 	dp = func(idx, prev int) int {
// 		if idx == n {
// 			return 0
// 		}
// 		if memo[idx][prev] != nil {
// 			return *memo[idx][prev]
// 		}
// 		// skip the current number
// 		best := dp(idx+1, prev)
// 		// keep/increase the current number the same
// 		for d := range 2 {
// 			nums[idx] += d
// 			if prev == -1 || prev+1 == nums[idx] {
// 				best = max(best, 1+dp(idx+1, nums[idx]))
// 			}
// 			nums[idx] -= d
// 		}
// 		memo[idx][prev] = &best
// 		return best
// 	}
// 	return dp(0, -1)
// }
