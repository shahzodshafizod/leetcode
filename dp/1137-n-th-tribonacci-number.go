package dp

// https://leetcode.com/problems/n-th-tribonacci-number/

// Tabulation
func tribonacci(n int) int {
	var cache [38]int

	cache[1], cache[2] = 1, 1
	for idx := 3; idx <= n; idx++ {
		cache[idx] = cache[idx-3] + cache[idx-2] + cache[idx-1]
	}

	return cache[n]
}

// // Memoization
// func tribonacci(n int) int {
// 	var cache [38]int
// 	var dp func(n int) int
// 	dp = func(n int) int {
// 		if n <= 1 {
// 			return n
// 		}
// 		if n == 2 {
// 			return 1
// 		}
// 		if cache[n] != 0 {
// 			return cache[n]
// 		}
// 		cache[n] = dp(n-3) + dp(n-2) + dp(n-1)
// 		return cache[n]
// 	}
// 	return dp(n)
// }
