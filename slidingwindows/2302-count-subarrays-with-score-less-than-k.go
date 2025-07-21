package slidingwindows

// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/

// Approach: Sliding Window
// Time: O(n)
// Space: O(1)
func countSubarrays2302(nums []int, k int64) int64 {
	var count, sum int64 = 0, 0
	start := 0
	for end := range nums {
		sum += int64(nums[end])
		for sum*int64(end-start+1) >= k {
			sum -= int64(nums[start])
			start++
		}
		count += int64(end - start + 1)
	}
	return count
}
