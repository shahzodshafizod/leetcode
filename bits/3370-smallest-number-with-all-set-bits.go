package bits

import "math/bits"

// https://leetcode.com/problems/smallest-number-with-all-set-bits/

func smallestNumber(n int) int {
	return (1 << bits.Len(uint(n))) - 1
}
