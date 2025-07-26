package greedy

// https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/

func maxSum(nums []int) int {
	total, maximal := 0, -((1 << 32) - 1)
	exists := make(map[int]bool)
	for _, num := range nums {
		if num > 0 {
			if !exists[num] {
				total += num
				exists[num] = true
			}
		} else if num > maximal {
			maximal = num
		}
	}
	if len(exists) > 0 {
		return total
	}
	return maximal
}
