package bits

// https://leetcode.com/problems/binary-number-with-alternating-bits/

// Approach: Optimized - Bit Manipulation with XOR trick
// If bits alternate: XOR with (n >> 1) gives all 1s
// Then (all_ones + 1) should be power of 2
// Example: n=5 (101) -> 101 ^ 010 = 111 -> 111 + 1 = 1000 (power of 2)
// Time: O(1) - constant time bit operations
// Space: O(1) - only variables
func hasAlternatingBits(n int) bool {
	// XOR with right-shifted version
	xor := n ^ (n >> 1)
	// If alternating, xor should be all 1s (like 111, 11111, etc.)
	// all_ones + 1 should be power of 2 (1000, 100000, etc.)
	// Power of 2 check: (x & (x+1)) == 0
	return (xor & (xor + 1)) == 0
}
