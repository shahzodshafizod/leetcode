package greedy

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/

func minDifference(nums []int) int {
	n := len(nums)
	if n <= 4 {
		return 0
	}

	sort.Ints(nums)

	minDiff := math.MaxInt

	for left := range 4 { // smaller index
		right := n - 4 + left // larger index
		minDiff = min(minDiff, nums[right]-nums[left])
	}

	return minDiff
}
