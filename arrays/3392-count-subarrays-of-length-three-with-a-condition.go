package arrays

// https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/

func countSubarrays(nums []int) int {
	count := 0
	for i, n := 2, len(nums); i < n; i++ {
		if (nums[i-2]+nums[i])*2 == nums[i-1] {
			count++
		}
	}
	return count
}
