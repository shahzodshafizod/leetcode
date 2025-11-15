package hashes

import "sort"

// https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	sort.Ints(nums)

	return nums[len(nums)/2]
}
