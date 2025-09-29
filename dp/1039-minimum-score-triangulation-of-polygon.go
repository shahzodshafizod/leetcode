package dp

import "math"

// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/

// Approach #2: Bottom-Up Dynamic Programming
// Time: O(nnn)
// Space: O(nn)
func minScoreTriangulation(values []int) int {
	n := len(values)

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}

	for j := 2; j < n; j++ {
		for i := j - 2; i >= 0; i-- {
			dp[i][j] = math.MaxInt
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j],
					dp[i][k]+values[i]*values[k]*values[j]+dp[k][j],
				)
			}
		}
	}

	return dp[0][n-1]
}

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(nnn)
// // Space: O(nn)
// func minScoreTriangulation(values []int) int {
// 	n := len(values)

// 	memo := make([][]*int, n)
// 	for idx := range memo {
// 		memo[idx] = make([]*int, n)
// 	}

// 	var dp func(i int, j int) int

// 	dp = func(i int, j int) int {
// 		if i+2 > j {
// 			return 0
// 		}

// 		if i+2 == j {
// 			return values[i] * values[i+1] * values[i+2]
// 		}

// 		if memo[i][j] != nil {
// 			return *memo[i][j]
// 		}

// 		res := math.MaxInt
// 		for k := i + 1; k < j; k++ {
// 			res = min(res, values[i]*values[k]*values[j]+dp(i, k)+dp(k, j))
// 		}

// 		memo[i][j] = &res

// 		return res
// 	}

// 	return dp(0, n-1)
// }
