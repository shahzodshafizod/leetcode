package twopointers

// https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	slow := 0

	for fast := range nums {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
