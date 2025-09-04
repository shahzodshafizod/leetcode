package maths

import "math"

// https://leetcode.com/problems/find-closest-person/

func findClosest(x int, y int, z int) int {
	dxz := int(math.Abs(float64(z - x)))
	dyz := int(math.Abs(float64(z - y)))

	if dxz < dyz {
		return 1
	} else if dyz < dxz {
		return 2
	}

	return 0
}
