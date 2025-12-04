package hashes

import "math"

// https://leetcode.com/problems/third-maximum-number/

// Approach: Set
// Time: O(n)
// Space: O(n)
func thirdMax(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	minimal := math.MinInt

	max1, max2, max3 := minimal, minimal, minimal
	for num := range set {
		if num > max1 {
			max3, max2, max1 = max2, max1, num
		} else if num > max2 {
			max3, max2 = max2, num
		} else if num > max3 {
			max3 = num
		}
	}

	if max3 == minimal {
		return max1
	}

	return max3
}
