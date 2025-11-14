package dp

import "math"

// https://leetcode.com/problems/tallest-billboard/

// Approach #3: Bottom-Up Dynamic Programming
// Time: O(n * total)
// Space: O(n * total)
func tallestBillboard(rods []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	n, total := len(rods), 0
	for _, rod := range rods {
		total += rod
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, total+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt
		}
	}

	dp[0][0] = 0

	for i := range n {
		for height := 0; height <= total; height++ {
			if dp[i][height] < 0 {
				continue
			}

			// Skip current rod
			dp[i+1][height] = max(dp[i+1][height], dp[i][height])

			// Include in left support
			newHeight := height + rods[i]
			if newHeight <= total {
				dp[i+1][newHeight] = max(dp[i+1][newHeight], dp[i][height])
			}

			// Include in right support
			newHeight = abs(height - rods[i])
			if newHeight <= total {
				dp[i+1][newHeight] = max(dp[i+1][newHeight], dp[i][height]+min(rods[i], height))
			}
		}
	}

	return max(0, dp[n][0])
}

// // Approach #2: Top-Down Dynamic Programming
// // Time: O(n * total)
// // Space: O(n * total)
// func tallestBillboard(rods []int) int {
// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}

// 		return x
// 	}

// 	n, total := len(rods), 0
// 	for _, rod := range rods {
// 		total += rod
// 	}

// 	memo := make([][]*int, n)
// 	for i := range memo {
// 		memo[i] = make([]*int, total+1)
// 	}

// 	var dp func(i int, height int) int

// 	dp = func(i int, height int) int {
// 		if i == n {
// 			if height == 0 {
// 				return 0
// 			}

// 			return math.MinInt
// 		}

// 		if 2*height > total {
// 			return math.MinInt
// 		}

// 		if memo[i][height] != nil {
// 			return *memo[i][height]
// 		}

// 		memo[i][height] = new(int)
// 		*memo[i][height] = max(
// 			dp(i+1, height), // skip
// 			dp(i+1, abs(height-rods[i]))+min(rods[i], height), // include in left
// 			dp(i+1, height+rods[i]),                           // include in right
// 		)

// 		return *memo[i][height]
// 	}

// 	return max(0, dp(0, 0))
// }

// // Approach #1: Brute-Force (Top-Down Dynamic Programming)
// // Time: O(3^(n/2))
// // Space: O(3^(n/2))
// func tallestBillboard(rods []int) int {
// 	n, total := len(rods), 0
// 	for _, rod := range rods {
// 		total += rod
// 	}

// 	memo := make([][][]*int, n)
// 	for i := range memo {
// 		memo[i] = make([][]*int, total+1)
// 		for j := range memo[i] {
// 			memo[i][j] = make([]*int, total+1)
// 		}
// 	}

// 	var dp func(i int, left int, right int) int

// 	dp = func(i int, left int, right int) int {
// 		if i == n {
// 			if left == right {
// 				return left
// 			}

// 			return 0
// 		}

// 		if memo[i][left][right] != nil {
// 			return *memo[i][left][right]
// 		}

// 		memo[i][left][right] = new(int)
// 		*memo[i][left][right] = max(
// 			dp(i+1, left, right),         // skip
// 			dp(i+1, left+rods[i], right), // include in left
// 			dp(i+1, left, right+rods[i]), // include in right
// 		)

// 		return *memo[i][left][right]
// 	}

// 	return dp(0, 0, 0)
// }
