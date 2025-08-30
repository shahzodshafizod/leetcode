package dp

// https://leetcode.com/problems/maximum-value-of-k-coins-from-piles/

// Approach: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(nkp), n=len(piles), p=len(piles[i])
// Space: O(k)
func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	curr, next := make([]int, k+1), make([]int, k+1)
	curr[1] = 1

	var total, limit int

	for i := n - 1; i >= 0; i-- {
		curr, next = next, curr
		for coins := 1; coins <= k; coins++ {
			// skip curr pile
			curr[coins] = next[coins]
			// take from curr pile
			total, limit = 0, min(coins, len(piles[i]))
			for j := range limit {
				total += piles[i][j]
				// skip curr pile
				curr[coins] = max(curr[coins],
					total+next[coins-j-1],
				)
			}
		}
	}

	return curr[k]
}

// // Approach: Bottom-Up Dynamic Programming
// // Time: O(nkp), n=len(piles), p=len(piles[i])
// // Space: O(nk)
// func maxValueOfCoins(piles [][]int, k int) int {
// 	var n = len(piles)
// 	var dp = make([][]int, n+1)
// 	for idx := range dp {
// 		dp[idx] = make([]int, k+1)
// 	}
// 	dp[n][1] = 1
// 	var total, limit int
// 	for i := n - 1; i >= 0; i-- {
// 		for coins := 1; coins <= k; coins++ {
// 			// skip curr pile
// 			dp[i][coins] = dp[i+1][coins]
// 			// take from curr pile
// 			total, limit = 0, min(coins, len(piles[i]))
// 			for j := 0; j < limit; j++ {
// 				total += piles[i][j]
// 				// skip curr pile
// 				dp[i][coins] = max(dp[i][coins],
// 					total+dp[i+1][coins-j-1],
// 				)
// 			}
// 		}
// 	}
// 	return dp[0][k]
// }

// // Approach: Top-Down Dynamic Programming
// // Time: O(nkp), p=len(piles[i])
// // Space: O(nk)
// func maxValueOfCoins(piles [][]int, k int) int {
// 	var n = len(piles)
// 	var memo = make([][]*int, n)
// 	for idx := range memo {
// 		memo[idx] = make([]*int, k+1)
// 	}
// 	var dp func(i int, coins int) int
// 	dp = func(i int, coins int) int {
// 		if i == n {
// 			return 0
// 		}
// 		if memo[i][coins] != nil {
// 			return *memo[i][coins]
// 		}
// 		// skip curr pile
// 		var count = dp(i+1, coins)
// 		// take from curr pile
// 		var total, limit = 0, min(coins, len(piles[i]))
// 		for j := 0; j < limit; j++ { // O(len(piles[i]))
// 			total += piles[i][j]
// 			// skip curr pile
// 			count = max(count, total+dp(i+1, coins-1-j))
// 		}
// 		memo[i][coins] = &count
// 		return count
// 	}
// 	return dp(0, k)
// }
