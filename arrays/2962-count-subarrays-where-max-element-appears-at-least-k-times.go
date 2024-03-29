package arrays

/*
Total number of subarrays (n is len(nums)):
n * (n + 1) / 2
*/

// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/

func countSubarrays2962(nums []int, k int) int64 {
	var maximum int
	var count int64 = 0
	var indexes []int
	for idx, num := range nums {
		if maximum == 0 || num > maximum {
			maximum = num
			indexes = []int{idx}
			count = 0
		} else if num == maximum {
			indexes = append(indexes, idx)
		}
		if len := len(indexes); len >= k {
			count += int64(indexes[len-k]) + 1
		}
	}
	return count
}
