package binarysearch

import "slices"

// https://leetcode.com/problems/house-robber-iv/

func minCapability(nums []int, k int) int {
	n := len(nums)
	low, high := 1, slices.Max(nums)
	var mid, count, idx int
	for low < high {
		mid = low + (high-low)/2
		count = 0
		for idx = 0; idx < n; idx++ {
			if nums[idx] <= mid {
				count++
				idx++
			}
		}
		if count >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
