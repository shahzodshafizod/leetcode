package greedy

import "sort"

// https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/

// Approach: Sort + Greedy
// Time: O(nlogn)
// Space: O(1)
func partitionArray(nums []int, k int) int {
	sort.Ints(nums)

	partitions, minimum := 1, nums[0]
	for _, maximum := range nums {
		if maximum-minimum > k {
			minimum = maximum
			partitions++
		}
	}

	return partitions
}
