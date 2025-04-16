package slidingwindows

// https://leetcode.com/problems/count-the-number-of-good-subarrays/

func countGood(nums []int, k int) int64 {
	var count = make(map[int]int)
	var goods int64 = 0
	var pairs, n = 0, len(nums)
	for start, end := 0, 0; end < n; end++ {
		pairs += count[nums[end]]
		count[nums[end]]++
		for pairs >= k {
			count[nums[start]]--
			pairs -= count[nums[start]]
			start++
			goods += int64(n - end)
		}
	}
	return goods
}
