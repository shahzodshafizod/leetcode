package slidingwindows

import "sort"

// https://leetcode.com/problems/find-k-th-smallest-pair-distance/

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	var countPairDistances = func(distance int) int {
		// Sliding Window
		// count total number of pairs
		// with diff <= distance
		var left, count = 0, 0
		for right := range nums {
			for left < right && nums[right]-nums[left] > distance {
				left++
			}
			count += right - left
		}
		return count
	}
	var left, right = 0, nums[len(nums)-1]
	var distance int
	for left < right {
		distance = left + (right-left)/2
		if countPairDistances(distance) >= k {
			right = distance
		} else {
			left = distance + 1
		}
	}
	return right
}
