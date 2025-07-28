package arrays

// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/

func longestMonotonicSubarray(nums []int) int {
	length := 1
	inc, dec := 1, 1

	for idx := len(nums) - 1; idx > 0; idx-- {
		if nums[idx-1] < nums[idx] {
			inc++
			dec = 1
		} else if nums[idx-1] > nums[idx] {
			dec++
			inc = 1
		} else {
			inc, dec = 1, 1
		}

		length = max(length, inc, dec)
	}

	return length
}
