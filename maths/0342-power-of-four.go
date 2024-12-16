package maths

// https://leetcode.com/problems/power-of-four/

// Time: O(1)
func isPowerOfFour(n int) bool {
	// All power of 4 numbers have 4 common features:
	// 1. greater than 0.
	var one = n > 0
	// 2. have one set bit in their binary representation.
	var two = n&(n-1) == 0
	// 3. the set bit is located in odd position from right to left.
	// The maximum number in binary representation with a length of 32 is:
	// bin(0b01010101010101010101010101010101) = hex(0x55555555) = dec(1431655765).
	var three = n&0b01010101010101010101010101010101 != 0
	// 4. subtracting 1 makes the number divisable to 3.
	var four = (n-1)%3 == 0
	return one && two && three && four
}

// // Approach #2: Iterative
// func isPowerOfFour(n int) bool {
// 	for n > 0 && n%4 == 0 {
// 		n /= 4
// 	}
// 	return n == 1
// }

// // Approach #1: Recursive
// func isPowerOfFour(n int) bool {
// 	if n <= 1 {
// 		return n == 1
// 	}
// 	return n%4 == 0 && isPowerOfFour(n/4)
// }
