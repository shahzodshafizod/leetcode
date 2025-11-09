package greedy

import (
	"sort"
)

// https://leetcode.com/problems/reducing-dishes/

// Approach #2: Greedy from the end
// Time: O(n log n)
// Space: O(1)
func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)

	maxsum, presum := 0, 0

	for i := len(satisfaction) - 1; i >= 0; i-- {
		presum += satisfaction[i]
		if presum < 0 {
			break
		}

		maxsum += presum
	}

	return maxsum
}

// // Approach #1: Brute-Force
// // Time: O(nn)
// // Space: O(1)
// func maxSatisfaction(satisfaction []int) int {
// 	sort.Ints(satisfaction)
// 	n := len(satisfaction)

// 	var maxsum, currSum, time int
// 	for start := range n {
// 		currSum, time = 0, 1
// 		for i := start; i < n; i++ {
// 			currSum += time * satisfaction[i]
// 			time++
// 		}

// 		maxsum = max(maxsum, currSum)
// 	}

// 	return maxsum
// }
