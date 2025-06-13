package binarysearch

import "sort"

// https://leetcode.com/problems/minimize-the-maximum-difference-of-pairs/

// Approach: Greedy + Binary Search
// Time: O(nlogn + nlogm), n=len(nums), m=max(nums)
// Space: O(1)
func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	var n = len(nums)
	var left, right = 0, nums[n-1]
	var mid, cnt int
	for left < right {
		mid = left + (right-left)/2

		cnt = 0
		for idx := 1; idx < n && cnt < p; idx++ {
			if nums[idx]-nums[idx-1] <= mid {
				cnt++
				idx++
			}
		}

		if cnt == p {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
