package maths

import "math"

// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/

func areaOfMaxDiagonal(dimensions [][]int) int {
	var (
		diag, d                float64
		area, length, width, a int
	)

	for idx := range dimensions {
		length = dimensions[idx][0]
		width = dimensions[idx][1]
		d = math.Sqrt(float64(length*length + width*width))

		a = length * width
		if d > diag || d == diag && a > area {
			diag = d
			area = a
		}
	}

	return area
}
