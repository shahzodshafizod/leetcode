package arrays

// https://leetcode.com/problems/1-bit-and-2-bit-characters/

// Approach #2: Greedy
// Time: O(n)
// Space: O(1)
func isOneBitCharacter(bits []int) bool {
	i, n := 0, len(bits)
	for i < n-1 {
		i += bits[i] + 1
	}

	return i == n-1
}

// // Approach #1: Single Pass Greedy
// // Time: O(n)
// // Space: O(1)
// func isOneBitCharacter(bits []int) bool {
// 	i, n := 0, len(bits)
// 	for i < n-1 { // Stop before last bit
// 		if bits[i] == 1 { // skip two bits: Two-bit character (10 or 11)
// 			i += 2
// 		} else { // skip one bit: One-bit character (0)
// 			i++
// 		}
// 	}

// 	return i == n-1
// }
