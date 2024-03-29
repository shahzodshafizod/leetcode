package hashes

// https://leetcode.com/problems/subarrays-with-k-different-integers/

func subarraysWithKDistinct(nums []int, k int) int {
	var atMost = func(k int) int {
		var count = make(map[int]int)
		var result = 0
		var start = 0
		for end := range nums {
			count[nums[end]]++
			for len(count) > k {
				count[nums[start]]--
				if count[nums[start]] == 0 {
					delete(count, nums[start])
				}
				start++
			}
			result += end - start + 1
		}
		return result
	}
	// exactly(k) = atMost(k) - atMost(k-1)
	return atMost(k) - atMost(k-1)
}
