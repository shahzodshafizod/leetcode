package binarysearch

// https://leetcode.com/problems/sqrtx/description/

func mySqrt(x int) int {
	var left, right = 1, x
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
