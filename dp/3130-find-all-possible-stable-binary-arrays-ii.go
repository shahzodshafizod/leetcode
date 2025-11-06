package dp

// https://leetcode.com/problems/find-all-possible-stable-binary-arrays-ii/

// Approach #3: Bottom-Up Dynamic Programming
// Time: O(zero * one)
// Space: O(zero * one)
func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = int(1e9) + 7

	dp := make([][][2]int, zero+1)
	for i := range dp {
		dp[i] = make([][2]int, one+1)
	}

	for z := min(zero, limit); z >= 0; z-- {
		dp[z][0][0] = 1
	}

	for o := min(one, limit); o >= 0; o-- {
		dp[0][o][1] = 1
	}

	for z := 1; z <= zero; z++ {
		for o := 1; o <= one; o++ {
			dp[z][o][0] = (dp[z-1][o][0] + dp[z-1][o][1]) % mod

			dp[z][o][1] = (dp[z][o-1][0] + dp[z][o-1][1]) % mod
			if z > limit {
				dp[z][o][0] = (dp[z][o][0] - dp[z-limit-1][o][1] + mod) % mod
			}

			if o > limit {
				dp[z][o][1] = (dp[z][o][1] - dp[z][o-limit-1][0] + mod) % mod
			}
		}
	}

	return (dp[zero][one][0] + dp[zero][one][1]) % mod
}

// // Approach #2: Bottom-Up Dynamic Programming (TLE + MLE)
// // Time: O(zero * one * limit)
// // Space: O(zero * one * limit)
// func numberOfStableArrays(zero int, one int, limit int) int {
// 	const mod = int(1e9) + 7
// 	dp := make([][][2]int, zero+1)
// 	for i := range dp {
// 		dp[i] = make([][2]int, one+1)
// 	}
// 	dp[0][0][0] = 1
// 	dp[0][0][1] = 1
// 	for z := range zero + 1 {
// 		for o := range one + 1 {
// 			if z == 0 && o == 0 {
// 				continue
// 			}
// 			for i := 1; i <= min(z, limit); i++ {
// 				dp[z][o][1] = (dp[z][o][1] + dp[z-i][o][0]) % mod
// 			}
// 			for i := 1; i <= min(o, limit); i++ {
// 				dp[z][o][0] = (dp[z][o][0] + dp[z][o-i][1]) % mod
// 			}
// 		}
// 	}
// 	return (dp[zero][one][0] + dp[zero][one][1]) % mod
// }

// // Approach #1: Top-Down Dynamic Programming (TLE + MLE)
// // Time: O(zero * one * limit)
// // Space: O(zero * one * limit)
// func numberOfStableArrays(zero int, one int, limit int) int {
// 	const mod = int(1e9) + 7
// 	memo := make([][][2]*int, zero+1)
// 	for z := range zero + 1 {
// 		memo[z] = make([][2]*int, one+1)
// 	}
// 	var dp func(zero int, one int, prev int) int
// 	dp = func(zero int, one int, prev int) int {
// 		if zero == 0 && one == 0 {
// 			return 1
// 		}
// 		if memo[zero][one][prev] != nil {
// 			return *memo[zero][one][prev]
// 		}
// 		var total int
// 		if prev == 1 {
// 			for z := min(zero, limit); z > 0; z-- {
// 				total = (total + dp(zero-z, one, 0)) % mod
// 			}
// 		} else {
// 			for o := min(one, limit); o > 0; o-- {
// 				total = (total + dp(zero, one-o, 1)) % mod
// 			}
// 		}
// 		memo[zero][one][prev] = &total
// 		return total
// 	}
// 	return (dp(zero, one, 0) + dp(zero, one, 1)) % mod
// }
