package arrays

// https://leetcode.com/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) []int {
	var n = len(nums)
	var newNums = make([]int, n)
	var lsq = nums[0] * nums[0]                   // left square
	var rsq = nums[n-1] * nums[n-1]               // right square
	for li, ri, i := 0, n-1, n-1; li <= ri; i-- { // left/right/newNums index
		switch {
		case lsq == -1:
			lsq = nums[li] * nums[li]
		case rsq == -1:
			rsq = nums[ri] * nums[ri]
		}
		if lsq >= rsq {
			newNums[i] = lsq
			li++
			lsq = -1
		} else {
			newNums[i] = rsq
			ri--
			rsq = -1
		}
	}
	return newNums
}
