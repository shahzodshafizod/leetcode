package bits

// https://leetcode.com/problems/power-of-two/

// Approach #3: Bit Manipulation
// Time: O(1)
// Space: O(1)
func isPowerOfTwo(n int) bool {
	// return bits.OnesCount(uint(n)) == 1
	return n > 0 && (n&(n-1)) == 0
}

// // Approach #2: Iterative
// // Time: O(logn)
// // Space: O(1)
// func isPowerOfTwo(n int) bool {
// 	var count = 0
// 	for n > 0 {
// 		count += n & 1
// 		n >>= 1
// 	}
// 	return count == 1
// }

// // Approach #1: Recursive
// // Time: O(logn)
// // Space: O(logn)
// func isPowerOfTwo(n int) bool {
// 	if n <= 1 {
// 		return n == 1
// 	}
// 	return n%2 == 0 && isPowerOfTwo(n/2)
// }
