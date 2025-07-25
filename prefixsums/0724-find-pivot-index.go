package prefixsums

// https://leetcode.com/problems/find-pivot-index/

func pivotIndex(nums []int) int {
	rightSum := 0
	for _, num := range nums {
		rightSum += num
	}
	leftSum := 0
	for idx, num := range nums {
		rightSum -= num
		if leftSum == rightSum {
			return idx
		}
		leftSum += num
	}
	return -1
}
