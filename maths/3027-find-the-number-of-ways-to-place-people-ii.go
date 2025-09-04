package maths

import "sort"

// https://leetcode.com/problems/find-the-number-of-ways-to-place-people-ii/

func numberOfPairsII(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}

		return points[i][0] < points[j][0]
	})
	n := len(points)

	var count, mny, mxy, y int
	for i := range n - 1 {
		mny, mxy = -1e9-1, points[i][1]
		for j := i + 1; j < n; j++ {
			y = points[j][1]
			if mny < y && y <= mxy {
				count++
				mny = y
			}
		}
	}

	return count
}
