package strings

// https://leetcode.com/problems/total-characters-in-string-after-transformations-i/

// Approach: Doubly-Linked List
func lengthAfterTransformations(s string, t int) int {
	var dp [26]int
	for _, c := range s {
		dp[int(c-'a')]++
	}

	const MOD int = 1e9 + 7

	z, a := 25, 0
	for ; t > 0; t-- {
		a = (z + 1) % 26
		dp[a] = (dp[a] + dp[z]) % MOD
		// y becomes z, so z moves back circly
		z = (z - 1 + 26) % 26
	}

	length := 0
	for _, cnt := range dp {
		length = (length + cnt) % MOD
	}

	return length
}

// func lengthAfterTransformations(s string, t int) int {
// 	var dp [26]int
// 	for _, c := range s {
// 		dp[int(c-'a')]++
// 	}
// 	const MOD int = 1e9 + 7
// 	var countz int
// 	for ; t > 0; t-- {
// 		countz = dp[25]
// 		for i := 25; i > 0; i-- {
// 			dp[i] = dp[i-1]
// 		}
// 		dp[0] = countz
// 		dp[1] = (dp[1] + countz) % MOD
// 	}
// 	var length = 0
// 	for _, cnt := range dp {
// 		length = (length + cnt) % MOD
// 	}
// 	return length
// }
