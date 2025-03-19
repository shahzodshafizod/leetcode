package slidingwindows

// https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/

func minOperations3191(nums []int) int {
	var flip, n = 0, len(nums)
	var curr, next = nums[0], nums[1]
	var prev int
	for idx := 2; idx < n; idx++ {
		prev, curr = curr, next
		next = nums[idx]
		if prev == 0 {
			curr ^= 1
			next ^= 1
			flip++
		}
	}
	if curr == 0 || next == 0 {
		flip = -1
	}
	return flip
}
