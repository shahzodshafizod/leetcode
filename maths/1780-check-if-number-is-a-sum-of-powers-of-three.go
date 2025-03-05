package maths

// https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/

// Approach #3: Ternary Representation
// Time: O( log3(n) )
// Space: O(1)
func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		}
		n /= 3
	}
	return true
}

// // Approach #2: Iterative (Optimized)
// // Time: O( log3(n) )
// // Space: O(1)
// func checkPowersOfThree(n int) bool {
// 	var power = 1
// 	for power*3 <= n {
// 		power *= 3
// 	}
// 	for n > 0 && power > 0 {
// 		if n >= power {
// 			n -= power
// 		}
// 		power /= 3
// 	}
// 	return n == 0
// }

// // Approach #1: Backtracking
// // Time: O(2 ^ log3(n))
// // Space: O( log3(n) )
// func checkPowersOfThree(n int) bool {
// 	var backtrack func(power int, sum int) bool
// 	backtrack = func(power int, sum int) bool {
// 		if sum == n {
// 			return true
// 		}
// 		if sum > n || power > n {
// 			return false
// 		}
// 		return backtrack(power*3, sum+power) ||
// 			backtrack(power*3, sum)
// 	}
// 	return backtrack(1, 0)
// }
