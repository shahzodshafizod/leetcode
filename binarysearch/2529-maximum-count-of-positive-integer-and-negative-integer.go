package binarysearch

// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/

func maximumCount(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] >= 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	neg := right + 1
	left, right = 0, n-1
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] <= 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	pos := n - left
	return max(neg, pos)
}
