package dp

// https://leetcode.com/problems/student-attendance-record-ii/

func checkRecord(n int) int {
	if n == 1 {
		return 3
	}

	const MOD = int(1e9 + 7)

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2
	dp[2] = 4

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % MOD
	}

	withA := 0
	for i := 1; i <= n; i++ {
		withA = (withA + dp[i-1]*dp[n-i]) % MOD
	}

	return (dp[n] + withA) % MOD
}

// func checkRecord(n int) int {
// 	const (
// 		MOD               = 1_000_000_007
// 		TOTAL_ABSENCES    = 1
// 		CONSECUTIVE_LATES = 2
// 	)
// 	var dp = make([][TOTAL_ABSENCES + 1][CONSECUTIVE_LATES + 1]int, n+1)
// 	dp[0] = [2][3]int{{1, 1, 1}, {1, 1, 1}}
// 	for length := 1; length <= n; length++ {
// 		for absences := 0; absences <= TOTAL_ABSENCES; absences++ {
// 			for lates := 0; lates <= CONSECUTIVE_LATES; lates++ {
// 				var val = dp[length-1][absences][2] // ...P
// 				if absences > 0 {
// 					val += dp[length-1][absences-1][2] // ...A
// 				}
// 				if lates > 0 {
// 					val += dp[length-1][absences][lates-1] // ...L
// 				}
// 				dp[length][absences][lates] = val % MOD
// 			}
// 		}
// 	}
// 	return dp[n][TOTAL_ABSENCES][CONSECUTIVE_LATES]
// }

// // time: O(n x 2 x 3) = O(6n) = O(n)
// // space: O(n x 2 x 3) = O(6n) = O(n)
// func checkRecord(n int) int {
// 	const (
// 		MOD               = 1_000_000_007
// 		TOTAL_ABSENCES    = 1
// 		CONSECUTIVE_LATES = 2
// 	)
// 	var cache = make([][TOTAL_ABSENCES + 1][CONSECUTIVE_LATES + 1]int, n+1) // [P][A][L]
// 	for i := 0; i <= n; i++ {
// 		for j := 0; j <= TOTAL_ABSENCES; j++ {
// 			for k := 0; k <= CONSECUTIVE_LATES; k++ {
// 				cache[i][j][k] = -1
// 			}
// 		}
// 	}
// 	var dfs func(records int, absences int, lates int) int
// 	dfs = func(records int, absences int, lates int) int {
// 		if records == 0 {
// 			return 1
// 		}
// 		if cache[records][absences][lates] != -1 {
// 			return cache[records][absences][lates]
// 		}
// 		var count = dfs(records-1, absences, CONSECUTIVE_LATES) // P
// 		if absences > 0 {
// 			count += dfs(records-1, absences-1, CONSECUTIVE_LATES) // A
// 		}
// 		if lates > 0 {
// 			count += dfs(records-1, absences, lates-1) // L
// 		}
// 		count %= MOD
// 		cache[records][absences][lates] = count
// 		return count
// 	}
// 	return dfs(n, TOTAL_ABSENCES, CONSECUTIVE_LATES)
// }
