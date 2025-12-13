package arrays

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

// Approach #2: Mark Visited with Negation
// Time: O(n) - two passes through array
// Space: O(1) - modify array in-place
func findDisappearedNumbers(nums []int) []int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}

		return x
	}

	// Mark visited numbers by negating values at their indices
	var idx int
	for _, num := range nums {
		idx = abs(num) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	// Collect indices with positive values
	var absent []int

	for idx := range nums {
		if nums[idx] > 0 {
			absent = append(absent, idx+1)
		}
	}

	return absent
}

// // Approach #1: Brute Force - Use Hash Set
// // Time: O(n) - create set and check each number
// // Space: O(n) - hash set
// func findDisappearedNumbers(nums []int) []int {
// 	numSet := make(map[int]bool)
// 	for _, num := range nums {
// 		numSet[num] = true
// 	}

// 	result := []int{}

// 	for i := 1; i <= len(nums); i++ {
// 		if !numSet[i] {
// 			result = append(result, i)
// 		}
// 	}

// 	return result
// }
