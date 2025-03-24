package intervals

import (
	"sort"
)

// https://leetcode.com/problems/count-days-without-meetings/

// Approach: Sorting
// Time: O(NLogN)
// Space: O(1)
func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0] ||
			meetings[i][1] < meetings[j][1]
	})
	var last = 0
	var start, end int
	for idx := range meetings {
		start, end = meetings[idx][0], meetings[idx][1]
		start = max(start, last+1)
		if end >= start {
			days -= end - start + 1
		}
		last = max(last, end)
	}
	return days
}
