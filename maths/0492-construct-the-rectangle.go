package maths

import "math"

// https://leetcode.com/problems/construct-the-rectangle/

// Approach: Find factors starting from sqrt
// Time: O(sqrt(n)) - check factors from sqrt down to 1
// Space: O(1) - constant space
func constructRectangle(area int) []int {
	width := int(math.Sqrt(float64(area)))
	for area%width != 0 {
		width--
	}

	length := area / width

	return []int{length, width}
}
