package binarysearch

// https://leetcode.com/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
