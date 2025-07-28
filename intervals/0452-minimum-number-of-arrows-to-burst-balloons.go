package intervals

import "sort"

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/

// time: O(n log n)
// space: O(n)
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })

	count := 0

	var prev []int
	for idx, curr := range points {
		if idx == 0 || prev[1] < curr[0] {
			count++
			prev = curr
		} else { // merging, keeps overlapping area
			prev[0] = curr[0]
			prev[1] = min(prev[1], curr[1])
		}
	}

	return count
}
