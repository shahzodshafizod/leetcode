package arrays

import "sort"

// https://leetcode.com/problems/array-partition/

// Approach: Sort and sum odd indices
// Sort array, then pair adjacent elements. Sum of mins is sum of elements at even indices.
// Time: O(n log n) - sorting
// Space: O(1) - ignoring sort space
func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}

	return sum
}
