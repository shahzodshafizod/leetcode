package arrays

import "sort"

// https://leetcode.com/problems/sort-the-jumbled-numbers/

// Time: O(N Log N)
// Space: O(N)
func sortJumbled(mapping []int, nums []int) []int {
	n := len(nums)
	mapped := make([]int, n)
	indices := make([]int, n)
	var places int
	for idx, num := range nums {
		indices[idx] = idx
		places = 1
		if num == 0 {
			mapped[idx] = mapping[0]
		}
		for num > 0 {
			mapped[idx] += mapping[num%10] * places
			num /= 10
			places *= 10
		}
	}
	sort.Slice(indices,
		func(i, j int) bool {
			i, j = indices[i], indices[j]
			if mapped[i] == mapped[j] {
				return i < j
			}
			return mapped[i] < mapped[j]
		},
	)
	sorted := make([]int, n)
	for idx := 0; idx < n; idx++ {
		sorted[idx] = nums[indices[idx]]
	}
	return sorted
}
