package slidingwindows

// https://leetcode.com/problems/maximum-erasure-value/

func maximumUniqueSubarray(nums []int) int {
	score, maxscore, left := 0, 0, 0
	count := make(map[int]int)
	for _, num := range nums {
		score += num
		count[num]++
		for count[num] > 1 {
			count[nums[left]]--
			score -= nums[left]
			left++
		}
		maxscore = max(maxscore, score)
	}
	return maxscore
}
