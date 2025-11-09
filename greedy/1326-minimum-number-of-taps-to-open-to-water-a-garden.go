package greedy

// https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/

// Same as: https://leetcode.com/problems/video-stitching/

// Approach #2: Greedy (Jump Game II)
// Time: O(n)
// Space: O(n)
func minTaps(n int, ranges []int) int {
	maxReach := make([]int, n+1)

	for i, r := range ranges {
		left := max(0, i-r)
		right := min(n, i+r)
		maxReach[left] = max(maxReach[left], right)
	}

	taps := 0
	currentEnd := 0
	farthest := 0

	// Greedy: try to reach position n
	for i := 0; i <= n; i++ {
		// If we can't reach this position, impossible
		if i > farthest {
			return -1
		}

		// Update the farthest we can reach
		farthest = max(farthest, maxReach[i])

		// If we've reached the end of current tap's range
		if i == currentEnd {
			taps++
			currentEnd = farthest

			// If we can already reach the end, we're done
			if currentEnd >= n {
				return taps
			}
		}
	}

	return -1
}

// // Approach #1: Brute-Force (Bottom-Up Dynamic Programming)
// // Time: O(nr), r=ranges[i]
// // Space: O(n)
// func minTaps(n int, ranges []int) int {
// 	maximum := n + 2

// 	dp := make([]int, n+1)

// 	for i := range dp {
// 		dp[i] = maximum
// 	}

// 	dp[0] = 0

// 	var left, right int
// 	for i, tapRange := range ranges {
// 		left = max(0, i-tapRange)
// 		right = min(n, i+tapRange)

// 		for j := left; j <= right; j++ {
// 			dp[j] = min(dp[j], dp[left]+1)
// 		}
// 	}

// 	if dp[n] == maximum {
// 		return -1
// 	}

// 	return dp[n]
// }
