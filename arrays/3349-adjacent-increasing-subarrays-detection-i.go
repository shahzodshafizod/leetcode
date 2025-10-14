package arrays

// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	prev, curr := 1, 1

	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] {
			curr++
		} else {
			prev, curr = curr, 1
		}

		// cur/2 means: Both subarrays are part of the same increasing segment
		if max(curr/2, min(prev, curr)) >= k {
			return true
		}
	}

	return false
}
