package arrays

import (
	"sort"
)

// https://leetcode.com/problems/magnetic-force-between-two-balls/

// time: O(NLogN+NLogMax)
// space:  O(1)
func maxDistance(position []int, m int) int {
	sort.Ints(position) // O(N Log N)
	n := len(position)
	max := position[n-1]
	low, high := 1, max/(m-1)
	var force, balls, prevPosition, curr int
	distance := 0
	for low <= high { // O(Log Max)
		force = (low + high) / 2
		balls = 1
		prevPosition = position[0]
		for curr = 1; curr < n; curr++ { // O(N)
			if position[curr]-prevPosition >= force {
				balls++
				prevPosition = position[curr]
			}
			if balls == m {
				break
			}
		}
		if balls == m {
			distance = force
			low = force + 1
		} else {
			high = force - 1
		}
	}
	return distance
}
