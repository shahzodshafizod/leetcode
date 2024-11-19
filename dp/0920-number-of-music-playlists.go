package dp

// https://leetcode.com/problems/number-of-music-playlists/

// Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
// Time: O(gn), g=goal
// Space: O(n)
func numMusicPlaylists(n int, goal int, k int) int {
	const MOD int = 1e9 + 7
	var prev, curr = make([]int, n+1), make([]int, n+1)
	curr[0] = 1
	for slots := 1; slots <= goal; slots++ {
		prev, curr = curr, prev
		curr[0] = 0
		for songs := 1; songs <= min(n, slots); songs++ {
			// add a new song: n-songs+1 = # of choices
			curr[songs] = (n - songs + 1) * prev[songs-1] % MOD
			// replay an old song: songs-k = # of choices
			if songs > k {
				curr[songs] = (curr[songs] + (songs-k)*prev[songs]) % MOD
			}
		}
	}
	return curr[n]
}

// // Approach #2: Bottom-Up Dynamic Programming
// // Time: O(gn), g=goal
// // Space: O(gn)
// func numMusicPlaylists(n int, goal int, k int) int {
// 	const MOD int = 1e9 + 7
// 	var dp = make([][]int, goal+1)
// 	for idx := range dp {
// 		dp[idx] = make([]int, n+1)
// 	}
// 	dp[0][0] = 1
// 	for slots := 1; slots <= goal; slots++ {
// 		for songs := 1; songs <= min(n, slots); songs++ {
// 			// add a new song: n-songs+1 = # of choices
// 			dp[slots][songs] = (n - songs + 1) * dp[slots-1][songs-1] % MOD
// 			// replay an old song: songs-k = # of choices
// 			if songs > k {
// 				dp[slots][songs] = (dp[slots][songs] +
// 					(songs-k)*dp[slots-1][songs]) % MOD
// 			}
// 		}
// 	}
// 	return dp[goal][n]
// }

// // Approach #1: Top-Down Dynamic Programming
// // Time: O(gn), g=goal
// // Space: O(gn)
// func numMusicPlaylists(n int, goal int, k int) int {
// 	const MOD int = 1e9 + 7
// 	var memo = make([][]*int, goal+1)
// 	for idx := range memo {
// 		memo[idx] = make([]*int, n+1)
// 	}
// 	var dp func(slots int, songs int) int
// 	dp = func(slots int, songs int) int {
// 		if slots == 0 && songs == 0 {
// 			return 1
// 		}
// 		if slots == 0 || songs == 0 {
// 			return 0
// 		}
// 		if memo[slots][songs] != nil {
// 			return *memo[slots][songs]
// 		}
// 		// add a new song: n-songs+1 = # of choices
// 		var count = (n - songs + 1) * dp(slots-1, songs-1) % MOD
// 		// replay an old song: songs-k = # of choices
// 		if songs > k {
// 			count = (count + (songs-k)*dp(slots-1, songs)) % MOD
// 		}
// 		memo[slots][songs] = &count
// 		return count
// 	}
// 	return dp(goal, n)
// }
