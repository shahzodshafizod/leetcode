package greedy

import (
	"sort"
)

// https://leetcode.com/problems/minimum-time-to-complete-all-tasks/

// Approach: Greedy Algorithm with Sorting
// Time Complexity: O(n*log(n) + n*m) where n=len(tasks), m=max(end)
// Space Complexity: O(m) where m=max(end) for the 'on' array
func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})

	var total, start, end, duration int

	on := make([]bool, 2001)

	for _, task := range tasks {
		start, end, duration = task[0], task[1], task[2]
		for time := start; time <= end; time++ {
			if on[time] {
				duration--
			}
		}

		for time := end; time >= start && duration > 0; time-- {
			if !on[time] {
				on[time] = true
				duration--
				total++
			}
		}
	}

	return total
}
