package dp

// https://leetcode.com/problems/stone-game-iv/

// Approach: Bottom-Up Dynamic Programming
// Time: O(n*sqrt(n))
// Space: O(n)
func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	dp[0] = false
	for curr := 1; curr <= n; curr++ {
		for x := 1; x*x <= curr; x++ {
			if !dp[curr-x*x] {
				dp[curr] = true
				break
			}
		}
	}
	return dp[n]
}

// // Approach: Top-Down Dynamic Programming
// // Time: O(n*sqrt(n))
// // Space: O(n*sqrt(n))
// func winnerSquareGame(n int) bool {
// 	var seen = make([]*bool, n+1)
// 	var dp func(curr int) bool
// 	dp = func(curr int) bool {
// 		if curr == 0 {
// 			return false
// 		}
// 		if seen[curr] != nil {
// 			return *seen[curr]
// 		}
// 		seen[curr] = new(bool)
// 		for x := 1; x*x <= curr; x++ {
// 			// if next player loses, the current one wins
// 			if !dp(curr - x*x) {
// 				*seen[curr] = true
// 				return true
// 			}
// 		}
// 		*seen[curr] = false
// 		return false
// 	}
// 	return dp(n)
// }
