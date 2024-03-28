package arrays

// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/

func countSubarrays(nums []int, k int64) int64 {
	var count, sum int64 = 0, 0
	var start = 0
	var length int64 = 0
	for end := range nums {
		sum += int64(nums[end])
		length++
		for sum*length >= k && start <= end {
			sum -= int64(nums[start])
			length--
			start++
		}
		count += length
	}
	return count
}
