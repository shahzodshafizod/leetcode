package maths

import (
	"strconv"
	"strings"
)

// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/

func getNoZeroIntegers(n int) []int {
	var b int
	for a := n / 2; a > 0; a-- {
		b = n - a
		if !strings.Contains(strconv.Itoa(a), "0") &&
			!strings.Contains(strconv.Itoa(b), "0") {
			return []int{a, b}
		}
	}

	return []int{-1, -1}
}
