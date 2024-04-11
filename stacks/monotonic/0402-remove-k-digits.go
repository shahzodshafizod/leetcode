package monotonic

import "bytes"

// https://leetcode.com/problems/remove-k-digits/

// time: O(n)
// space: O(n)
func removeKdigits(num string, k int) string {
	var size = 0
	var stack = make([]byte, size, len(num)-k)
	var curr byte
	for idx := range num {
		curr = num[idx]
		for k > 0 && size > 0 && stack[size-1] > curr {
			stack = stack[:size-1]
			size--
			k--
		}
		stack = append(stack, curr)
		size++
	}
	num = string(bytes.TrimLeft(stack[:size-k], "0"))
	if num == "" {
		num = "0"
	}
	return num
}
