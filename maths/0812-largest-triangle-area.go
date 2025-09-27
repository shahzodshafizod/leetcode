package maths

import "math"

// https://leetcode.com/problems/largest-triangle-area/

/*
Heron's formula:
Area = √(s(s-a)(s-b)(s-c))
s = (a+b+c)/2
a = √(x2-x1)^2 + (y2-y1)^2)
*/

/*
Shoelace formula for a triangle
Area = ½ |(x₁y₂ + x₂y₃ + x₃y₁) - (y₁x₂ + y₂x₃ + y₃x₁)|
*/

func largestTriangleArea(points [][]int) float64 {
	n := len(points)

	var (
		x1, y1, x2, y2, x3, y3 int
		area                   float64
	)

	for i := range n {
		x1, y1 = points[i][0], points[i][1]
		for j := 1; j < n-1; j++ {
			x2, y2 = points[j][0], points[j][1]
			for k := j + 1; k < n; k++ {
				x3, y3 = points[k][0], points[k][1]
				area = max(
					area,
					0.5*math.Abs(float64(x1*y2+x2*y3+x3*y1)-float64(y1*x2+y2*x3+y3*x1)),
				)
			}
		}
	}

	return area
}
