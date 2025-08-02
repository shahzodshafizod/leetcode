package maths

// https://leetcode.com/problems/plus-one/

func plusOne(digits []int) []int {
	for idx := len(digits) - 1; idx >= 0; idx-- {
		if digits[idx] < 9 {
			digits[idx]++

			return digits
		}

		digits[idx] = 0
	}

	digits = append(digits, 0)
	copy(digits, digits[1:])
	digits[0] = 1

	return digits
}
