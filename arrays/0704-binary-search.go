package arrays

// https://leetcode.com/problems/binary-search/

func search(nums []int, target int) int {
	len := len(nums)
	left, right := 0, len-1
	if target < nums[left] || target > nums[right] {
		return -1
	}
	var mid int
	for left <= right {
		mid = (left + right) >> 1 // "x>>1" == "x/2"
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
