package dp

import (
	"math/bits"
)

// https://leetcode.com/problems/count-k-reducible-numbers-less-than-n/

// // Approach 1: Brute Force (Only works for small values of n)
// // Time Complexity: O(n * k * log(n)) - iterates n numbers, each takes k reductions
// // Space Complexity: O(1)
// func countKReducibleNumbers(s string, k int) int {
// 	isKReducible := func(num int64, k int) bool {
// 		steps := 0
// 		for num > 1 && steps < k {
// 			num = int64(bits.OnesCount(uint(num)))
// 			steps++
// 		}

// 		return num == 1
// 	}

// 	const mod int = 1e9 + 7

// 	// Convert binary string to big integer
// 	n := new(big.Int)
// 	n.SetString(s, 2)

// 	var count int

// 	for i, nInt := int64(1), n.Int64(); i < nInt; i++ {
// 		if isKReducible(i, k) {
// 			count = (count + 1) % mod
// 		}
// 	}

// 	return count
// }

// // Approach 2: Top-Down DP (Memoization) with Digit DP
// // Time Complexity: O(2nn) where n is length of string
// // Space Complexity: O(2nn) for memoization
// func countKReducibleNumbers(s string, k int) int {
// 	const mod int = 1e9 + 7

// 	n := len(s)

// 	steps := make([]int, n+1)
// 	for i := 2; i <= n; i++ {
// 		steps[i] = steps[bits.OnesCount(uint(i))] + 1
// 	}

// 	memo := make([][][2]*int, n)
// 	for i := range memo {
// 		memo[i] = make([][2]*int, n)
// 	}

// 	var dp func(pos int, ones int, bound int) int

// 	dp = func(pos int, ones int, bound int) int {
// 		if pos == n {
// 			if bound == 0 && ones > 0 && steps[ones] < k {
// 				return 1
// 			}

// 			return 0
// 		}

// 		if memo[pos][ones][bound] != nil {
// 			return *memo[pos][ones][bound]
// 		}

// 		limit := 1
// 		if bound == 1 {
// 			limit = int(s[pos] - '0')
// 		}

// 		result := 0

// 		for digit := 0; digit <= limit; digit++ {
// 			newBound := bound
// 			if digit != limit {
// 				newBound = 0
// 			}

// 			result = (result + dp(pos+1, ones+digit, newBound)) % mod
// 		}

// 		memo[pos][ones][bound] = &result

// 		return result
// 	}

// 	return dp(0, 0, 1)
// }

// Approach 3: Bottom-Up DP (Tabulation)
// Time Complexity: O(2nn) where n is length of string
// Space Complexity: O(2nn)
func countKReducibleNumbers(s string, k int) int {
	const mod int = 1e9 + 7

	n := len(s)

	steps := make([]int, n+1)
	for i := 2; i <= n; i++ {
		steps[i] = steps[bits.OnesCount(uint(i))] + 1
	}

	dp := make([][][2]int, n+1)
	for i := range n + 1 {
		dp[i] = make([][2]int, n+1)
	}

	for ones := range n + 1 {
		if ones > 0 && steps[ones] < k {
			dp[n][ones][0] = 1
		}
	}

	for pos := n - 1; pos >= 0; pos-- {
		for ones := range n + 1 {
			for bound := range 2 {
				limit := 1
				if bound == 1 {
					limit = int(s[pos] - '0')
				}

				result := 0

				for digit := 0; digit <= limit && ones+digit <= n; digit++ {
					newOnes := ones + digit
					newBound := bound

					if digit != limit {
						newBound = 0
					}

					result = (result + dp[pos+1][newOnes][newBound]) % mod
				}

				dp[pos][ones][bound] = result
			}
		}
	}

	return dp[0][0][1]
}
