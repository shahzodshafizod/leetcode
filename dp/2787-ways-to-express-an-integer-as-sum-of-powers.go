package dp

import "math"

// https://leetcode.com/problems/ways-to-express-an-integer-as-sum-of-powers/

// Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(nn)
// Space: O(n)
func numberOfWays(n int, x int) int {
	const mod = int(1e9) + 7

	dp := make([]int, n+1)
	dp[0] = 1

	var power int
	for i := 1; i <= n; i++ {
		power = int(math.Pow(float64(i), float64(x)))
		for j := n; j >= power; j-- {
			dp[j] = (dp[j] + dp[j-power]) % mod
		}
	}

	return dp[n]
}

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(nn)
// // Space: O(nn)
// func numberOfWays(n int, x int) int {
// 	const mod = int(1e9) + 7

// 	dp := make([][]int, n+1)
// 	for idx := range n + 1 {
// 		dp[idx] = make([]int, n+1)
// 	}

// 	dp[0][0] = 1

// 	var power int
// 	for i := 1; i <= n; i++ {
// 		power = int(math.Pow(float64(i), float64(x)))
// 		for j := range n + 1 {
// 			dp[i][j] = dp[i-1][j] // not take
// 			if j >= power {       // take
// 				dp[i][j] = (dp[i][j] + dp[i-1][j-power]) % mod
// 			}
// 		}
// 	}

// 	return dp[n][n]
// }

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(nn)
// // Space: O(nn)
// func numberOfWays(n int, x int) int {
// 	const mod = int(1e9) + 7

// 	memo := make([][]*int, n+1)
// 	for idx := range n + 1 {
// 		memo[idx] = make([]*int, n+1)
// 	}

// 	var dp func(k int, n int) int

// 	dp = func(k, n int) int {
// 		if n == 0 {
// 			return 1
// 		}

// 		if n < 0 || k == 0 {
// 			return 0
// 		}

// 		if memo[k][n] != nil {
// 			return *memo[k][n]
// 		}

// 		ways := dp(k-1, n)

// 		degree := int(math.Pow(float64(k), float64(x)))
// 		if n >= degree {
// 			ways = (ways + dp(k-1, n-degree)) % mod
// 		}

// 		memo[k][n] = &ways

// 		return ways
// 	}

// 	return dp(n, n)
// }
