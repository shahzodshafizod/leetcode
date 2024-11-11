package slidingwindows

import "math"

// https://leetcode.com/problems/maximum-average-subarray-i/

func findMaxAverage(nums []int, k int) float64 {
	var total, maxTotal float64 = 0, -math.MaxFloat64
	for end := range nums {
		total += float64(nums[end])
		if end+1 >= k {
			maxTotal = max(maxTotal, total)
			total -= float64(nums[end+1-k])
		}
	}
	return maxTotal / float64(k)
}
