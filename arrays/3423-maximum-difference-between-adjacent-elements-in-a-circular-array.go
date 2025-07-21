package arrays

// https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/

func maxAdjacentDistance(nums []int) int {
	diff := -100
	n := len(nums)
	nei := n - 1
	for idx := 0; idx < n; idx++ {
		if nums[idx] > nums[nei] {
			diff = max(diff, nums[idx]-nums[nei])
		} else {
			diff = max(diff, nums[nei]-nums[idx])
		}
		nei = idx
	}
	return diff
}
