package hashes

import "math"

// https://leetcode.com/problems/max-points-on-a-line/

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	count := 0

	var x1, y1, x2, y2 int

	var slope float64

	for i := range points {
		x1, y1 = points[i][0], points[i][1]
		slopes := make(map[float64]int)

		for j := i + 1; j < n; j++ {
			x2, y2 = points[j][0], points[j][1]
			slope = math.MaxFloat64

			if x1 != x2 {
				// slope for any two points = slope= (y2-y1)/(x2-x1)
				slope = float64(y2-y1) / float64(x2-x1)
			}

			slopes[slope]++
			count = max(count, slopes[slope])
		}
	}

	return count + 1
}
