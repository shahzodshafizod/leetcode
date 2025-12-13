package greedy

import "sort"

// https://leetcode.com/problems/set-intersection-size-at-least-two/

// Approach: Greedy with Sorting
// Time: O(n log n) - dominated by sorting
// Space: O(1) - constant extra space (excluding sort)
func intersectionSizeTwo(intervals [][]int) int {
	// Sort by end point, then by start point descending
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] > intervals[j][0]
	})

	result := 0
	first, second := -1, -1

	for _, interval := range intervals {
		start, end := interval[0], interval[1]

		// Case 1: Both points already in interval
		if start <= first {
			continue
		}

		// Case 2: Only second point in interval
		if start <= second {
			first = second
			second = end
			result++
		} else {
			// Case 3: Neither point in interval
			first = end - 1
			second = end
			result += 2
		}
	}

	return result
}
