package arrays

import "sort"

// https://leetcode.com/problems/assign-cookies/

// Approach: Greedy with Sorting
// Sort both greed factors and cookie sizes.
// Assign smallest available cookie that satisfies each child.
// Time: O(n log n + m log m) - sorting
// Space: O(1) - excluding sort space
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	child := 0
	cookie := 0

	for child < len(g) && cookie < len(s) {
		if s[cookie] >= g[child] {
			child++
		}

		cookie++
	}

	return child
}
