package intervals

import (
	"sort"
)

// https://leetcode.com/problems/merge-intervals/

// time: O(n log n)
// space: O(n)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	var result = make([][]int, 0)
	var lastIdx = -1
	for _, interval := range intervals { // O(n)
		if lastIdx == -1 || result[lastIdx][1] < interval[0] {
			result = append(result, interval)
			lastIdx++
		} else {
			result[lastIdx][1] = max(result[lastIdx][1], interval[1])
		}
	}
	return result
}
