package arrays

import "math"

// https://leetcode.com/problems/minimize-manhattan-distances/

// Approach: Optimized - Transform coordinates using Chebyshev distance
// Manhattan dist = |x1-x2| + |y1-y2| = max(|(x1+y1)-(x2+y2)|, |(x1-y1)-(x2-y2)|)
// Transform to (u,v) = (x+y, x-y), max Manhattan = max(max(u)-min(u), max(v)-min(v))
// Time: O(n) for each candidate removal
// Space: O(n) for transformed coordinates
func minimizeMaxDistance(points [][]int) int {
	n := len(points)

	// Transform coordinates
	u := make([]int, n)
	v := make([]int, n)

	for i, p := range points {
		u[i] = p[0] + p[1]
		v[i] = p[0] - p[1]
	}

	getMaxDistWithout := func(skipIdx int) int {
		uMin, uMax := math.MaxInt32, math.MinInt32
		vMin, vMax := math.MaxInt32, math.MinInt32

		for i := range n {
			if i == skipIdx {
				continue
			}

			if u[i] < uMin {
				uMin = u[i]
			}

			if u[i] > uMax {
				uMax = u[i]
			}

			if v[i] < vMin {
				vMin = v[i]
			}

			if v[i] > vMax {
				vMax = v[i]
			}
		}

		uRange := uMax - uMin

		vRange := vMax - vMin
		if uRange > vRange {
			return uRange
		}

		return vRange
	}

	// Find extremes
	uMin, uMax := u[0], u[0]
	vMin, vMax := v[0], v[0]
	uMinIdx, uMaxIdx := 0, 0
	vMinIdx, vMaxIdx := 0, 0

	for i := 1; i < n; i++ {
		if u[i] < uMin {
			uMin = u[i]
			uMinIdx = i
		}

		if u[i] > uMax {
			uMax = u[i]
			uMaxIdx = i
		}

		if v[i] < vMin {
			vMin = v[i]
			vMinIdx = i
		}

		if v[i] > vMax {
			vMax = v[i]
			vMaxIdx = i
		}
	}

	// Try removing candidates
	candidates := []int{uMinIdx, uMaxIdx, vMinIdx, vMaxIdx}
	minResult := math.MaxInt32

	for _, skip := range candidates {
		maxDist := getMaxDistWithout(skip)
		if maxDist < minResult {
			minResult = maxDist
		}
	}

	return minResult
}
