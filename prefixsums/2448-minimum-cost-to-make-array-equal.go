package prefixsums

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-cost-to-make-array-equal/

// Approach: Sorting + Prefix Sum
// Time: O(nlogn + c), n=len(nums), c=max(cost[i])
// Space: O(n)
func minCost(nums []int, cost []int) int64 {
	n := len(nums)

	indices := make([]int, n)
	for idx := range indices {
		indices[idx] = idx
	}

	sort.Slice(indices, func(i, j int) bool { return nums[indices[i]] < nums[indices[j]] })

	var postMult, postCost int64 = 0, 0
	for idx := 0; idx < n; idx++ {
		postMult += int64(nums[idx] * cost[idx])
		postCost += int64(cost[idx])
	}

	var preMult, preCost int64 = 0, 0

	var minTotal int64 = math.MaxInt64

	for _, idx := range indices {
		preMult += int64(nums[idx] * cost[idx])
		postMult -= int64(nums[idx] * cost[idx])
		preCost += int64(cost[idx])
		postCost -= int64(cost[idx])
		total := int64(nums[idx])*preCost - preMult
		total += postMult - int64(nums[idx])*postCost
		minTotal = min(minTotal, total)
	}

	return minTotal
}

// // Approach: Brute-Force (Time Limit Exceeded)
// // Time: O(N*C), N=len(nums), C=max(cost[i])
// // Space: O(1)
// func minCost(nums []int, cost []int) int64 {
// 	var abs = func(n int) int {
// 		if n < 0 {
// 			return -n
// 		}
// 		return n
// 	}
// 	var minTotal int64 = math.MaxInt64
// 	for _, target := range nums {
// 		var total int64 = 0
// 		for idx := range nums {
// 			total += int64(abs(nums[idx]-target)) * int64(cost[idx])
// 		}
// 		minTotal = min(minTotal, total)
// 	}
// 	return minTotal
// }
