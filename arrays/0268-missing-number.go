package arrays

// https://leetcode.com/problems/missing-number/

func missingNumber(nums []int) int {
	var n = len(nums)
	var sum = n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
