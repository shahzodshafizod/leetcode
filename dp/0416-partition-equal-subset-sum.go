package dp

import "maps"

// https://leetcode.com/problems/partition-equal-subset-sum/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(n*sum)
// Space: O(sum)
func canPartition(nums []int) bool {
	var total = 0
	for _, num := range nums {
		total += num
	}
	if total&1 != 0 {
		return false
	}
	var target = total / 2
	var dp = map[int]bool{0: true}
	for _, num := range nums {
		var nextdp = maps.Clone(dp)
		for t := range dp {
			if t+num == target {
				return true
			}
			nextdp[t+num] = true
		}
		dp = nextdp
	}
	return dp[target]
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(n*sum)
// // Space: O(n*sum)
// func canPartition(nums []int) bool {
// 	var n = len(nums)
// 	var sum = 0
// 	for idx := 0; idx < n; idx++ {
// 		sum += nums[idx]
// 	}
// 	if sum&1 != 0 {
// 		return false
// 	}
// 	var half = sum / 2
// 	var memo = make([][]*bool, n)
// 	for idx := 0; idx < n; idx++ {
// 		memo[idx] = make([]*bool, half+1)
// 	}
// 	var dfs func(idx int, sum int) bool
// 	dfs = func(idx int, sum int) bool {
// 		if idx == n || sum <= 0 {
// 			return sum == 0
// 		}
// 		if memo[idx][sum] != nil {
// 			return *memo[idx][sum]
// 		}
// 		memo[idx][sum] = new(bool)
// 		*memo[idx][sum] = dfs(idx+1, sum) ||
// 			dfs(idx+1, sum-nums[idx])
// 		return *memo[idx][sum]
// 	}
// 	return dfs(0, half)
// }
