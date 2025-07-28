package greedy

import "sort"

// https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/

// Approach: Sorting
// Time: O(nlogn)
// Space: O(n)
func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)

	var parts [][]int

	for idx, n := 2, len(nums); idx < n; idx += 3 {
		if nums[idx]-nums[idx-2] > k {
			return nil
		}

		parts = append(parts, nums[idx-2:idx+1])
	}

	return parts
}
