package dp

import "sort"

// https://leetcode.com/problems/find-the-sum-of-subsequence-powers/

// Approach #2: Top-Down Dynamic Programming
// Time: O(nn k d), n=len(nums), d=diff
// Space: O(nn k d)
func sumOfPowers(nums []int, k int) int {
	sort.Ints(nums)
	const mod = int(1e9) + 7
	n := len(nums)
	memo := make(map[[3]int]*int)
	var dp func(idx int, k int, diff int) int
	dp = func(idx, k, diff int) int {
		if k == 0 {
			return diff
		}
		key := [3]int{idx, k, diff}
		if memo[key] != nil {
			return *memo[key]
		}
		var total int
		for nxt := idx + 1; nxt < n; nxt++ {
			total = (total + dp(nxt, k-1, min(diff, nums[nxt]-nums[idx]))) % mod
		}
		memo[key] = &total
		return total
	}
	inf := (1 << 32) - 1
	var total int
	for i := range n {
		total = (total + dp(i, k-1, inf)) % mod
	}
	return total
}

// // Approach #1: Top-Down Dynamic Programming (slower)
// // Time: O(nn k d), n=len(nums), d=diff
// // Space: O(nn k d)
// func sumOfPowers(nums []int, k int) int {
// 	sort.Ints(nums)
// 	const mod = int(1e9) + 7
// 	memo := make(map[[4]int]*int)
// 	var dp func(prev int, curr int, k int, diff int) int
// 	dp = func(prev, curr, k, diff int) int {
// 		if k == 0 {
// 			return diff
// 		}
// 		if curr == len(nums) {
// 			return 0
// 		}
// 		key := [4]int{prev, curr, k, diff}
// 		if memo[key] != nil {
// 			return *memo[key]
// 		}
// 		// skip
// 		total := dp(prev, curr+1, k, diff)
// 		// include
// 		if prev >= 0 {
// 			diff = min(diff, nums[curr]-nums[prev])
// 		}
// 		total = (total + dp(curr, curr+1, k-1, diff)) % mod
// 		memo[key] = &total
// 		return total
// 	}
// 	return dp(-1, 0, k, (1<<32)-1)
// }
