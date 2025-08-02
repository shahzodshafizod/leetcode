package bits

import "math/bits"

// https://leetcode.com/problems/find-the-k-th-character-in-string-game-ii/

// Approach: Mathematics
// Time: O(logk)
// Space: O(1)
func kthCharacter(k int64, operations []int) byte {
	ans := 0

	for n := bits.Len(uint(k)); n >= 0; n-- {
		if k > 1<<n {
			k -= 1 << n
			ans += operations[n]
		}
	}

	return byte('a' + ans%26)
}
