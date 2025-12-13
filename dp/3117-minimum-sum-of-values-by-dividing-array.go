package dp

import "math"

// https://leetcode.com/problems/minimum-sum-of-values-by-dividing-array/

// Approach: Dynamic Programming with memoization
// Divide array into m segments where each segment's bitwise AND equals andValues[j]
// Minimize sum of last elements of each segment
// Time: O(n * m * unique_and_values) where unique_and_values is bounded by log(max_val)
// Space: O(n * m * unique_and_values) for memoization
func minimumValueSum(nums []int, andValues []int) int {
	n := len(nums)
	m := len(andValues)
	memo := make(map[[3]int]int)

	var dp func(i, j, currentAnd int) int

	dp = func(i, j, currentAnd int) int {
		// Base cases
		if i == n && j == m {
			return 0
		}

		if i == n || j == m {
			return math.MaxInt32 / 2
		}

		state := [3]int{i, j, currentAnd}
		if val, exists := memo[state]; exists {
			return val
		}

		// Update AND with current element
		newAnd := currentAnd & nums[i]

		result := math.MaxInt32 / 2

		// Option 1: Continue current segment
		result = min(result, dp(i+1, j, newAnd))

		// Option 2: End current segment if AND matches
		if newAnd == andValues[j] {
			result = min(result, nums[i]+dp(i+1, j+1, (1<<20)-1))
		}

		memo[state] = result

		return result
	}

	// Start with all bits set (AND identity)
	ans := dp(0, 0, (1<<20)-1)
	if ans >= math.MaxInt32/2 {
		return -1
	}

	return ans
}
