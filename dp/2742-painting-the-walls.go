package dp

// https://leetcode.com/problems/painting-the-walls/

// Approach: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(nn)
// Space: O(n)
func paintWalls(cost []int, time []int) int {
	var n = len(cost)
	var dp = make([][]int, n+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}
	var curr, next = make([]int, n+1), make([]int, n+1)
	for walls := 1; walls <= n; walls++ {
		curr[walls] = 1e9
	}
	for idx := n - 1; idx >= 0; idx-- {
		curr, next = next, curr
		for wall := 1; wall <= n; wall++ {
			curr[wall] = min(
				cost[idx]+next[max(0, wall-1-time[idx])],
				next[wall],
			)
		}
	}
	return curr[n]
}

// // Approach: Bottom-Up Dynamic Programming
// // Time: O(nn)
// // Space: O(nn)
// func paintWalls(cost []int, time []int) int {
// 	var n = len(cost)
// 	var dp = make([][]int, n+1)
// 	for idx := range dp {
// 		dp[idx] = make([]int, n+1)
// 	}
// 	for walls := 1; walls <= n; walls++ {
// 		dp[n][walls] = 1e9
// 	}
// 	for idx := n - 1; idx >= 0; idx-- {
// 		for wall := 1; wall <= n; wall++ {
// 			dp[idx][wall] = min(
// 				cost[idx]+dp[idx+1][max(0, wall-1-time[idx])],
// 				dp[idx+1][wall],
// 			)
// 		}
// 	}
// 	return dp[0][n]
// }

// // Approach: Top-Down Dynamic Programming
// // Time: O(nn)
// // Space: O(nn)
// func paintWalls(cost []int, time []int) int {
// 	var n = len(cost)
// 	var memo = make([][]*int, n)
// 	for idx := range memo {
// 		memo[idx] = make([]*int, n+1)
// 	}
// 	var dp func(idx int, walls int) int
// 	dp = func(idx int, walls int) int {
// 		if walls <= 0 {
// 			return 0
// 		}
// 		if idx == n {
// 			return 1e9
// 		}
// 		if memo[idx][walls] != nil {
// 			return *memo[idx][walls]
// 		}
// 		memo[idx][walls] = new(int)
// 		*memo[idx][walls] = min(
// 			cost[idx]+dp(idx+1, walls-1-time[idx]),
// 			dp(idx+1, walls),
// 		)
// 		return *memo[idx][walls]
// 	}
// 	return dp(0, n)
// }
