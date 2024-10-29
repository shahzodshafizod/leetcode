package dp

import "math"

// https://leetcode.com/problems/stone-game-iii/

// Approach #1: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(n)
// Space: O(1)
func stoneGameIII(stoneValue []int) string {
	var n = len(stoneValue)
	var dp [3]int
	var best, stones int
	for idx := n - 1; idx >= 0; idx-- {
		best = math.MinInt
		stones = 0
		for x := idx; x < min(n, idx+3); x++ {
			stones += stoneValue[x]
			best = max(best, stones-dp[x-idx])
		}
		dp[0], dp[1], dp[2] = best, dp[0], dp[1]
	}
	if dp[0] > 0 {
		return "Alice"
	}
	if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(n)
// // Space: O(n)
// func stoneGameIII(stoneValue []int) string {
// 	var n = len(stoneValue)
// 	var dp = make([]int, n+1)
// 	var stones int
// 	for idx := n - 1; idx >= 0; idx-- {
// 		dp[idx] = math.MinInt
// 		stones = 0
// 		for x := idx; x < min(n, idx+3); x++ {
// 			stones += stoneValue[x]
// 			dp[idx] = max(dp[idx], stones-dp[x+1])
// 		}
// 	}
// 	if dp[0] > 0 {
// 		return "Alice"
// 	}
// 	if dp[0] < 0 {
// 		return "Bob"
// 	}
// 	return "Tie"
// }

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(n)
// // Space: O(n)
// func stoneGameIII(stoneValue []int) string {
// 	var n = len(stoneValue)
// 	var cache = make([]*int, n)
// 	var dp func(idx int) int
// 	dp = func(idx int) int {
// 		if idx == n {
// 			return 0
// 		}
// 		if cache[idx] != nil {
// 			return *cache[idx]
// 		}
// 		var total, stones = math.MinInt, 0
// 		for x := idx; x < min(n, idx+3); x++ {
// 			stones += stoneValue[x]
// 			total = max(total, stones-dp(x+1)) // 1-(2-3)=1-2+3
// 		}
// 		cache[idx] = &total
// 		return total
// 	}
// 	if dp(0) > 0 {
// 		return "Alice"
// 	}
// 	if dp(0) < 0 {
// 		return "Bob"
// 	}
// 	return "Tie"
// }
