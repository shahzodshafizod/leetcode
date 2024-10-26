package dp

// https://leetcode.com/problems/k-inverse-pairs-array/

/*
n=3, k=2
|   | 0 | 1 | 2 |
|---+---+---+---|
| 0 | 1 | 0 | 0 |
| 1 | 1 | 0 | 0 |
| 2 | 1 | 1 | 0 |
| 3 | 1 | 2 | 2 |
dp[n][k] - # of arrays with k inverse pairs, given n elements
*/

// Approach #3: Bottom-Up Dynamic Programming (Time & Space Optimized)
// Time: O(N * K)
// Space: O(K)
func kInversePairs(n int, k int) int {
	const MOD int = 1e9 + 7
	var prev = make([]int, k+1)
	var curr = make([]int, k+1)
	curr[0] = 1
	var total int
	for ni := 1; ni <= n; ni++ {
		prev, curr = curr, prev
		total = 0
		for ki := 0; ki <= k; ki++ {
			if ki >= ni {
				// (+ MOD) % MOD - to avoid -ve
				total = (total - prev[ki-ni] + MOD) % MOD
			}
			total = (total + prev[ki]) % MOD
			curr[ki] = total
		}
	}
	return curr[k]
}

// // Approach #2: Bottom-Up Dynamic Programming (Time Limit Exceeded)
// // Time: O(N^2 * K)
// // Space: O(N * K)
// func kInversePairs(n int, k int) int {
// 	const MOD int = 1e9 + 7
// 	var count = make([][]int, n+1)
// 	for idx := range count {
// 		count[idx] = make([]int, k+1)
// 	}
// 	count[0][0] = 1
// 	for ni := 1; ni <= n; ni++ {
// 		for ki := 0; ki <= k; ki++ {
// 			for pairs := 0; pairs < min(ni, ki+1); pairs++ {
// 				count[ni][ki] = (count[ni][ki] + count[ni-1][ki-pairs]) % MOD
// 			}
// 		}
// 	}
// 	return count[n][k]
// }

// // Approach #1: Top-Down Dynamic Programming (Time Limit Exceeded)
// // Time: O(N^2 * K)
// // Space: O(N * K)
// func kInversePairs(n int, k int) int {
// 	const MOD int = 1e9 + 7
// 	var dp = make([][]*int, n+1)
// 	for idx := range dp {
// 		dp[idx] = make([]*int, k+1)
// 	}
// 	var count func(n int, k int) int
// 	count = func(n int, k int) int {
// 		if k <= 0 {
// 			if k == 0 {
// 				return 1
// 			}
// 			return 0
// 		}
// 		if dp[n][k] != nil {
// 			return *dp[n][k]
// 		}
// 		dp[n][k] = new(int)
// 		for i := 0; i < n; i++ {
// 			*dp[n][k] = (*dp[n][k] + count(n-1, k-i)) % MOD
// 		}
// 		return *dp[n][k]
// 	}
// 	return count(n, k)
// }
