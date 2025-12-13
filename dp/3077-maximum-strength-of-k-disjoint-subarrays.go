package dp

// https://leetcode.com/problems/maximum-strength-of-k-disjoint-subarrays/

// Approach: Dynamic Programming (Kadane-like for each level)
// dp[i][j] = max strength selecting i subarrays from first j elements
// For each subarray level i, use Kadane's algorithm to find max weighted subarray sum
// multiplier for i-th subarray = (k - i + 1) if i is odd, -(k - i + 1) if i is even
// Time: O(n * k), Space: O(n * k)
func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	inf := int64(-1e18)

	// dp[i][j] = max strength selecting i subarrays from first j elements
	dp := make([][]int64, k+1)
	for i := range dp {
		dp[i] = make([]int64, n+1)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= k; i++ {
		// multiplier for i-th subarray: (k - i + 1) if i is odd, else -(k - i + 1)
		multiplier := int64(k - i + 1)
		if i%2 == 0 {
			multiplier = -multiplier
		}

		maxSum := inf
		current := inf

		for j := i - 1; j < n; j++ {
			// current = max sum ending at position j for subarray level i
			// Either extend previous subarray or start new from dp[i-1][j]
			val1 := current + int64(nums[j])*multiplier

			val2 := dp[i-1][j] + int64(nums[j])*multiplier
			if val1 > val2 {
				current = val1
			} else {
				current = val2
			}

			if current > maxSum {
				maxSum = current
			}

			dp[i][j+1] = maxSum
		}
	}

	if dp[k][n] > inf {
		return dp[k][n]
	}

	return 0
}
