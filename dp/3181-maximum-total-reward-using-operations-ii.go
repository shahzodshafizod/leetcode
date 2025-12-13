package dp

import (
	"math/big"
	"sort"
)

// https://leetcode.com/problems/maximum-total-reward-using-operations-ii/

// Approach #2: Dynamic Programming with Big Integer Bitmask
// Time: O(n log n + n * max_val) where max_val is maximum reward value
// Space: O(max_val) for bitmask
func maxTotalReward(rewardValues []int) int {
	removeDuplicatesAndSort := func(arr []int) []int {
		seen := make(map[int]bool)
		result := []int{}

		for _, val := range arr {
			if !seen[val] {
				seen[val] = true
				result = append(result, val)
			}
		}

		sort.Ints(result)

		return result
	}
	rewards := removeDuplicatesAndSort(rewardValues)

	// Bitmask: bit i is set if we can achieve sum i
	// Initially, we can achieve sum 0
	possible := big.NewInt(1) // binary: 00...001

	for _, reward := range rewards {
		// We can only use this reward if current sum < reward
		// Get all possible sums less than reward
		mask := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), uint(reward)), big.NewInt(1))
		smallerSums := new(big.Int).And(possible, mask)

		// Add reward to all smaller sums
		newSums := new(big.Int).Lsh(smallerSums, uint(reward))
		possible.Or(possible, newSums)
	}

	// Find the highest bit set in possible (maximum sum)
	return possible.BitLen() - 1
}

// // Approach #1: Brute Force Backtracking
// // Time: O(2^n) - exponential, TLE for large inputs
// // Space: O(n) - recursion depth
// func maxTotalReward(rewardValues []int) int {
// 	removeDuplicatesAndSort := func(arr []int) []int {
// 		seen := make(map[int]bool)
// 		result := []int{}

// 		for _, val := range arr {
// 			if !seen[val] {
// 				seen[val] = true
// 				result = append(result, val)
// 			}
// 		}

// 		sort.Ints(result)

// 		return result
// 	}

// 	rewards := removeDuplicatesAndSort(rewardValues)

// 	var backtrack func(index, currentSum int) int

// 	backtrack = func(index, currentSum int) int {
// 		if index == len(rewards) {
// 			return currentSum
// 		}

// 		// Skip current reward
// 		maxReward := backtrack(index+1, currentSum)

// 		// Take current reward if possible
// 		if currentSum < rewards[index] {
// 			takeReward := backtrack(index+1, currentSum+rewards[index])
// 			if takeReward > maxReward {
// 				maxReward = takeReward
// 			}
// 		}

// 		return maxReward
// 	}

// 	return backtrack(0, 0)
// }
