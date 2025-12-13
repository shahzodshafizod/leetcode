package dp

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-increments-for-target-multiples-in-an-array/

// Approach: Optimized DP with LCM Calculation
// Time: O(n * 3^m * m) where 3^m comes from subset iteration
// Space: O(2^m)
func minOperations(nums []int, target []int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}

		return a
	}

	lcm := func(a, b int) int {
		return a * b / gcd(a, b)
	}

	m := len(target)
	inf := math.MaxInt32
	size := 1 << m

	// Precompute LCM for each non-empty subset
	lcmMap := make(map[int]int)

	for mask := 1; mask < size; mask++ {
		lcmVal := 1

		for i := range m {
			if mask&(1<<i) != 0 {
				lcmVal = lcm(lcmVal, target[i])
			}
		}

		lcmMap[mask] = lcmVal
	}

	dp := make([]int, size)
	for i := range dp {
		dp[i] = inf
	}

	dp[0] = 0

	sort.Ints(nums)

	for _, num := range nums {
		newDp := make([]int, size)
		copy(newDp, dp)

		for mask := range size {
			if dp[mask] == inf {
				continue
			}

			uncovered := (size - 1) ^ mask
			for subset := uncovered; subset > 0; subset = (subset - 1) & uncovered {
				// Calculate exact cost using LCM
				lcmVal := lcmMap[subset]
				remainder := num % lcmVal

				cost := 0
				if remainder != 0 {
					cost = lcmVal - remainder
				}

				newMask := mask | subset
				if dp[mask]+cost < newDp[newMask] {
					newDp[newMask] = dp[mask] + cost
				}
			}
		}

		dp = newDp
	}

	if dp[size-1] == inf {
		return -1
	}

	return dp[size-1]
}
