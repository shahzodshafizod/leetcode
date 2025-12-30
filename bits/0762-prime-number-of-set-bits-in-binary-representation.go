package bits

import "math/bits"

// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/

// // Approach: Bit Counting with Prime Check Optimization
// // Key insight: Since right <= 10^6 < 2^20, the maximum number of set bits is 20
// // Prime numbers up to 20: {2, 3, 5, 7, 11, 13, 17, 19}
// //
// // Algorithm:
// // 1. For each number in [left, right], count the number of set bits
// // 2. Check if the count is a prime number using a precomputed bitmask
// // 3. Count how many numbers have a prime number of set bits
// //
// // Optimization: Use a bitmask to represent primes up to 20
// // primes = {2, 3, 5, 7, 11, 13, 17, 19}
// // bitmask = 0b10100010100010101100 (bits 2,3,5,7,11,13,17,19 are set)
// // bitmask = 665772 in decimal
// //
// // Time: O((right - left) * log(right)) - count bits for each number
// // Space: O(1) - only using constant space
// func countPrimeSetBits(left int, right int) int {
// 	// Primes up to 20: {2, 3, 5, 7, 11, 13, 17, 19}
// 	// Binary representation with these bits set: 0b10100010100010101100
// 	const primesMask = 0b10100010100010101100 // = 665772

// 	count := 0

// 	for num := left; num <= right; num++ {
// 		// Count set bits using Brian Kernighan's algorithm
// 		setBits := 0

// 		n := num
// 		for n > 0 {
// 			n &= n - 1 // Remove the rightmost set bit
// 			setBits++
// 		}

// 		// Check if setBits is prime using the bitmask
// 		if primesMask&(1<<setBits) != 0 {
// 			count++
// 		}
// 	}

// 	return count
// }

// Alternative: Using bits.OnesCount from math/bits
func countPrimeSetBits(left int, right int) int {
	const primesMask = 0b10100010100010101100
	count := 0
	for num := left; num <= right; num++ {
		setBits := bits.OnesCount(uint(num))
		if primesMask&(1<<setBits) != 0 {
			count++
		}
	}
	return count
}
