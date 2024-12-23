package bits

// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/

// Note: This question is the same as 461: Hamming Distance.

// Approach #3: Brian Kernighan’s Algorithm
// Time: O(number of set bits)
// Space: O(1)
func minBitFlips(start int, goal int) int {
	// return bits.OnesCount(uint(start ^ goal))
	var flips = 0
	var result = start ^ goal
	for result != 0 {
		result &= result - 1
		flips += 1
	}
	return flips
}

// // Approach #2: Iterative
// // Time: O(max bits)
// // Space: O(1)
// func minBitFlips(start int, goal int) int {
// 	var flips = 0
// 	var result = start ^ goal
// 	for result != 0 {
// 		flips += result & 1
// 		result >>= 1
// 	}
// 	return flips
// }

// // Approach #1: Recursive
// // Time: O(max bits)
// // Space: O(max bits)
// func minBitFlips(start int, goal int) int {
// 	if start == 0 && goal == 0 {
// 		return 0
// 	}
// 	var flips = (start & 1) ^ (goal & 1)
// 	return flips + minBitFlips(start>>1, goal>>1)
// }
