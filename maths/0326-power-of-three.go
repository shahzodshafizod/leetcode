package maths

// https://leetcode.com/problems/power-of-three/

// Time: O(log3(n))
// Space: O(1)
func isPowerOfThree(n int) bool {
	// if n <= 0 {
	// 	return false
	// }

	// for ; n%3 == 0; n /= 3 {
	// }

	// return n == 1

	// every power of 3 divides to the earliest powers of 3.
	// in 32 bit numbers, 3^20 is the largest power of three
	// 3^20 = 3486784401
	return n > 0 && 3486784401%n == 0
}
