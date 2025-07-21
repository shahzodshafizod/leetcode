package slidingwindows

// https://leetcode.com/problems/count-complete-subarrays-in-an-array/

func countCompleteSubarrays(nums []int) int {
	count := make(map[int]int)
	k := 0
	for _, num := range nums {
		if count[num] == 0 {
			k++
		}
		count[num] = -1
	}
	total, start := 0, 0
	for _, num := range nums {
		if count[num] == -1 {
			k--
		}
		count[num]++
		for k == 0 {
			count[nums[start]]--
			if count[nums[start]] == -1 {
				k++
			}
			start++
		}
		total += start
	}
	return total
}
