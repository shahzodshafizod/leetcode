package bits

// https://leetcode.com/problems/neighboring-bitwise-xor/

/*
Since the derived array was created by XORing the original elements,
and each element of the original array participates in the XORing
of precisely two elements of the derived array,
XORing all of the derived array's elements once more should result
in the output of disregarding every element of the original array.
That result should be equal to zero.
*/

// Approach: Bit Manipulation
// Time: O(n)
// Space: O(1)
func doesValidArrayExist(derived []int) bool {
	xorsum := 0
	for _, item := range derived {
		xorsum ^= item
	}

	return xorsum == 0
}
