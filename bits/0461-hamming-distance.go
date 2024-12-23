package bits

import "math/bits"

// https://leetcode.com/problems/hamming-distance/

// Note: This question is the same as 2220: Minimum Bit Flips to Convert Number.

// Approach #3: Bit Manipulation
// Time: O(1)
// Space: O(1)
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

// // Approach #2: Iterative
// // Time: O(b), b=# of bits of max(start, goal)
// // Space: O(1)
// func hammingDistance(x int, y int) int {
// 	var distance = 0
// 	for x > 0 || y > 0 {
// 		distance += (x & 1) ^ (y & 1)
// 		x >>= 1
// 		y >>= 1
// 	}
// 	return distance
// }

// // Approach #1: Recursive
// // Time: O(b), b=# of bits of max(start, goal)
// // Space: O(b)
// func hammingDistance(x int, y int) int {
// 	if x == 0 && y == 0 {
// 		return 0
// 	}
// 	return (x&1 ^ y&1) + hammingDistance(x>>1, y>>1)
// }
