package dp

// https://leetcode.com/problems/count-substrings-divisible-by-last-digit/

// Approach: Dynamic Programming with Space Optimization
// Key insight: For each position, track remainder of substrings ending here
// when divided by each possible last digit (1-9).
//
// Strategy:
// 1. Use dp[divisor][remainder] = count of substrings with this remainder
// 2. For each new digit, update all divisors (1-9)
// 3. Count when remainder is 0 and divisor equals current digit
// 4. Optimize space: rolling array (current and previous state)
//
// Mathematical Analysis:
//   - For substring with value V, extending with digit d:
//     New value = V * 10 + d
//     New remainder mod i = (old_remainder * 10 + d) % i
//
// Time Complexity: O(n * 9 * 9) = O(n)
// Space Complexity: O(9 * 9) = O(1)
func countSubstrings3448(s string) int64 {
	n := len(s)
	var result int64 = 0

	// dp[divisor][remainder] = count of substrings ending at current position
	// with this remainder when divided by divisor
	dp := make([][]int64, 10)
	for i := range dp {
		dp[i] = make([]int64, 10)
	}

	for idx := range n {
		digit := int(s[idx] - '0')

		// Create new state for current position
		newDp := make([][]int64, 10)
		for i := range newDp {
			newDp[i] = make([]int64, 10)
		}

		// For each possible divisor (1-9, can't divide by 0)
		for divisor := 1; divisor <= 9; divisor++ {
			// New substring starting at current position
			newDp[divisor][digit%divisor]++

			// Extend existing substrings
			for remainder := 0; remainder < divisor; remainder++ {
				if dp[divisor][remainder] > 0 {
					newRemainder := (remainder*10 + digit) % divisor
					newDp[divisor][newRemainder] += dp[divisor][remainder]
				}
			}
		}

		// Count substrings ending at current position divisible by last digit
		if digit != 0 {
			result += newDp[digit][0]
		}

		dp = newDp
	}

	return result
}

// // Approach: Brute Force - Check All Substrings
// // For each substring, track its value modulo the last digit and check divisibility.
// //
// // Time Complexity: O(n^2)
// // Space Complexity: O(1)
// func countSubstrings3448(s string) int64 {
// 	n := len(s)
// 	var count int64 = 0

// 	for i := range n {
// 		// For each possible divisor (last digit), track value mod that divisor
// 		mod := make([]int, 10)

// 		for j := i; j < n; j++ {
// 			digit := int(s[j] - '0')

// 			// Update modulo values for all divisors
// 			for divisor := 1; divisor <= 9; divisor++ {
// 				mod[divisor] = (mod[divisor]*10 + digit) % divisor
// 			}

// 			// Check if divisible by last digit
// 			if digit != 0 && mod[digit] == 0 {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }
