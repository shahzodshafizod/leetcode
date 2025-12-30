package arrays

// https://leetcode.com/problems/largest-number-at-least-twice-of-others/

// Approach 2: One-Pass Solution (Optimized)
// Track both the largest and second largest elements in one pass
// Then check if largest >= 2 * second_largest
// Time: O(n) - single pass through the array
// Space: O(1) - only using a few variables
func dominantIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	maxVal := -1
	maxIdx := -1
	secondMax := -1

	// Find both max and second max in one pass
	for i, num := range nums {
		if num > maxVal {
			secondMax = maxVal
			maxVal = num
			maxIdx = i
		} else if num > secondMax {
			secondMax = num
		}
	}

	// Check if max is at least twice the second max
	if maxVal >= 2*secondMax {
		return maxIdx
	}

	return -1
}

// // Approach 1: Two-Pass Solution
// // First pass: Find the largest element and its index
// // Second pass: Check if largest is at least twice every other element
// // Time: O(n) - two passes through the array
// // Space: O(1) - only using a few variables
// func dominantIndex(nums []int) int {
// 	// First pass: find max and its index
// 	maxVal := nums[0]
// 	maxIdx := 0

// 	for i, num := range nums {
// 		if num > maxVal {
// 			maxVal = num
// 			maxIdx = i
// 		}
// 	}

// 	// Second pass: check if max is at least twice every other element
// 	for _, num := range nums {
// 		if num != maxVal && maxVal < 2*num {
// 			return -1
// 		}
// 	}

// 	return maxIdx
// }
