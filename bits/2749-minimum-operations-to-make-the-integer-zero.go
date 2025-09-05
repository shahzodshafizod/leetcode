package bits

import "math/bits"

// https://leetcode.com/problems/minimum-operations-to-make-the-integer-zero/

func makeTheIntegerZero(num1 int, num2 int) int {
	// num1 = (num2 + 2^i1) + (nums2 + 2^i2) + ... + (num2 + 2^ik)
	// num1 = k*num2 + (2^i1 + 2^i2 + ... + 2^ik)
	// num1 - k*num2 = 2^i1 + 2^i2 + ... + 2^ik
	var target int
	for k := 1; k <= 61; k++ {
		target = num1 - k*num2
		// checking if target is equal to: 2^i1 + 2^i2 + ... + 2^ik
		if k <= target && bits.OnesCount(uint(target)) <= k {
			return k
		}
	}

	return -1
}
