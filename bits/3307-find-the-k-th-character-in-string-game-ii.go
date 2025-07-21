package bits

import "math/bits"

// https://leetcode.com/problems/find-the-k-th-character-in-string-game-ii/

// Approach: Mathematics
// Time: O(logk)
// Space: O(1)
func kthCharacter(k int64, operations []int) byte {
	ans := 0
	for bits := bits.Len(uint(k)); bits >= 0; bits-- {
		if k > 1<<bits {
			k -= 1 << bits
			ans += operations[bits]
		}
	}
	return byte('a' + ans%26)
}
