package twopointers

import "sort"

// https://leetcode.com/problems/count-the-number-of-fair-pairs/

// Approach: Two Pointers
// Time: O(n log n)
// Space: O(1)
func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	lessThan := func(top int) int64 {
		var pairs int64 = 0
		left, right := 0, len(nums)-1
		for left < right {
			if nums[left]+nums[right] < top {
				pairs += int64(right - left)
				left++
			} else {
				right--
			}
		}
		return pairs
	}
	return lessThan(upper+1) - lessThan(lower)
}
