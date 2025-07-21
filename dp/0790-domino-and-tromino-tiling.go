package dp

/*
Full[n] = Full[n-1] + Full[n-2] + TopMissing[n-1] + BottomMissing[n-1]

  F[n-1]      F[n-2]        T[n-1]       B[n-1]
|#######+|  |#######--|  |#######-+|  |########-|
|#######+|  |#######--|  |########-|  |#######-+|


TopMissing[n] = Full[n-2] + BottomMissing[n-1]
  F[n-2]        B[n-1]
|#######- |  |######## |
|#######+-|  |#######--|


BottomMissing[n] = Full[n-2] + TopMissing[n-1]
  F[n-2]        T[n-1]
|#######+-|  |#######--|
|#######- |  |######## |
*/

// https://leetcode.com/problems/domino-and-tromino-tiling/

// Approach: Bottom-Up Dynamic Programming
// Time: O(n)
// Space: O(1)
func numTilings(n int) int {
	tm, bm := 0, 0
	preprev, prev, curr := 1, 1, 1 // [n-2], [n-1], [n]
	const MOD int = 1e9 + 7
	for i := 2; i <= n; i++ {
		curr = (prev + preprev + tm + bm) % MOD
		tm, bm = (preprev+bm)%MOD, (preprev+tm)%MOD
		preprev, prev = prev, curr
	}
	return curr
}

// // Approach: Bottom-Up Dynamic Programming
// // Time: O(n)
// // Space: O(n)
// func numTilings(n int) int {
// 	var full = make([]int, n+1)
// 	var tm = make([]int, n+1)
// 	var bm = make([]int, n+1)
// 	full[0], full[1] = 1, 1
// 	const MOD int = 1e9 + 7
// 	for i := 2; i <= n; i++ {
// 		full[i] = (full[i-1] + full[i-2] + tm[i-1] + bm[i-1]) % MOD
// 		tm[i] = (full[i-2] + bm[i-1]) % MOD
// 		bm[i] = (full[i-2] + tm[i-1]) % MOD
// 	}
// 	return full[n]
// }
