package arrays

// https://leetcode.com/problems/maximum-difference-between-increasing-elements/

func maximumDifference(nums []int) int {
	var premin, diff = nums[0], -1
	for _, curr := range nums {
		if curr > premin {
			diff = max(diff, curr-premin)
		} else {
			premin = curr
		}
	}
	return diff
}
