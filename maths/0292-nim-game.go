package maths

// https://leetcode.com/problems/nim-game/

/*
n=1: Win (take 1 stone)
n=2: Win (take 2 stones)
n=3: Win (take 3 stones)
n=4: Lose! No matter what we take:
     - Take 1 → opponent faces 3 (wins)
     - Take 2 → opponent faces 2 (wins)
     - Take 3 → opponent faces 1 (wins)
n=5: Win (take 1, leave 4 for opponent)
n=6: Win (take 2, leave 4 for opponent)
n=7: Win (take 3, leave 4 for opponent)
n=8: Lose (same pattern as 4)
*/

// Approach #2: Use mathematical pattern recognition
// Time: O(1)
// Space: O(1)
func canWinNim(n int) bool {
	return n%4 != 0
}

// // Approach #1: Use DP to determine winning positions
// // Time: O(n)
// // Space: O(n) for the DP array
// func canWinNim(n int) bool {
// 	// dp[i] represents whether we can win with i stones
// 	dp := make([]bool, n+1)
// 	dp[1], dp[2], dp[3] = true, true, true

// 	for i := 4; i <= n; i++ {
// 		// We win if we can make a move that leaves opponent in losing position
// 		// We can take 1, 2, or 3 stones
// 		dp[i] = !dp[i-1] || !dp[i-2] || !dp[i-3]
// 	}

// 	return dp[n]
// }
