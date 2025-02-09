package hashes

// https://leetcode.com/problems/count-number-of-bad-pairs/

func countBadPairs(nums []int) int64 {
	var n = len(nums)
	var pairs = make(map[int]int)
	var goods int64 = 0
	for idx, num := range nums {
		goods += int64(pairs[num-idx])
		pairs[num-idx]++
	}
	var bads int64 = int64(n*(n-1))/2 - goods
	return bads
}
