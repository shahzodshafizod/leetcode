package arrays

// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/

func countSubarrays(nums []int, k int64) int64 {
	var count, sum int64 = 0, 0
	var start = 0
	var score int64
	for end, num := range nums {
		sum += int64(num)
		score = sum * int64(end-start+1)
		for score >= k && start < end {
			sum -= int64(nums[start])
			start++
			score = sum * int64(end-start+1)
		}
		if score < k {
			count += int64(end - start + 1)
		}
	}
	return count
}
