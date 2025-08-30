package dp

import (
	"math/bits"
)

// https://leetcode.com/problems/maximize-score-after-n-operations/

// Approach: Bottom-Up Dynamic Programming
// Time: O(n^2 + 2^n + log(max(nums)))
// Space: O(2^n)
func maxScore(nums []int) int {
	var gcd func(a int, b int) int

	gcd = func(a int, b int) int {
		if b == 0 {
			return a
		}

		return gcd(b, a%b)
	}
	n := len(nums)
	limit := 1 << n
	dp := make([]int, limit)

	var count, newMask int
	for mask := range limit {
		count = bits.OnesCount(uint(mask))
		if count&1 == 1 {
			continue
		}

		for i := range n {
			if mask&(1<<i) == 0 {
				continue
			}

			for j := i + 1; j < n; j++ {
				if mask&(1<<j) == 0 {
					continue
				}

				newMask = mask ^ (1 << i) ^ (1 << j)
				dp[mask] = max(
					dp[mask],
					count/2*gcd(nums[i], nums[j])+dp[newMask],
				)
			}
		}
	}

	return dp[limit-1]
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(n^2 + 2^n + log(max(nums)))
// // Space: O(2^n)
// func maxScore(nums []int) int {
// 	var gcd func(a int, b int) int // O(log(max(nums)))
// 	gcd = func(a int, b int) int {
// 		if b == 0 {
// 			return a
// 		}
// 		return gcd(b, a%b)
// 	}
// 	var n = len(nums)
// 	var memo = make([]*int, 1<<n)
// 	var dp func(mask int, op int) int // O(2^n)
// 	dp = func(mask int, op int) int {
// 		if memo[mask] != nil {
// 			return *memo[mask]
// 		}
// 		memo[mask] = new(int)
// 		var newMask int
// 		for i := 0; i < n; i++ { // O(n^2)
// 			for j := i + 1; j < n; j++ {
// 				if mask&(1<<i) != 0 || mask&(1<<j) != 0 {
// 					continue
// 				}
// 				newMask = mask | (1 << i) | (1 << j)
// 				*memo[mask] = max(
// 					*memo[mask],
// 					op*gcd(nums[i], nums[j])+dp(newMask, op+1),
// 				)
// 			}
// 		}
// 		return *memo[mask]
// 	}
// 	return dp(0, 1)
// }
