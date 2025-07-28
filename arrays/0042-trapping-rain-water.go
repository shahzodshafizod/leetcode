package arrays

/*
Problem:
Given an array of integers representing an elevation map where the width of each bar is 1,
return how much rainwater can be trapped.

Step 1: Verify the constraints
	- Do the left and right sides of the graph count as walls?
		: No, the sides are not walls.
	- Will there be negative integers?
		: No, assume all integers are positive.
Step 2: Write out some test cases
	- [0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2] -> 8
	- [] -> 0
	- [3] -> 0
	- [3, 4, 3] -> 0
	- [5, 0, 3, 0, 0, 0, 2, 3, 4, 2] -> ?
Step 3: Figure out a solution without code
Step 6: Test out code with out test cases
*/

// https://leetcode.com/problems/trapping-rain-water/

func trap(height []int) int {
	total := 0
	maxElem := 0
	left, right := 0, len(height)-1

	for left < right {
		oldMaxElem := maxElem

		if height[left] < height[right] {
			maxElem = max(maxElem, height[left])
			total += (right-left)*(maxElem-oldMaxElem) - height[left]
			left++
		} else {
			maxElem = max(maxElem, height[right])
			total += (right-left)*(maxElem-oldMaxElem) - height[right]
			right--
		}
	}

	return total
}

// func trap(height []int) int {
// 	var total = 0
// 	var left, right = 0, len(height) - 1
// 	var maxLeft, maxRight = 0, 0
// 	for left < right {
// 		if height[left] <= height[right] {
// 			if height[left] >= maxLeft {
// 				maxLeft = height[left]
// 			} else {
// 				total += maxLeft - height[left]
// 			}
// 			left++
// 		} else {
// 			if height[right] >= maxRight {
// 				maxRight = height[right]
// 			} else {
// 				total += maxRight - height[right]
// 			}
// 			right--
// 		}
// 	}
// 	return total
// }

// func trap(height []int) int {
// 	var length = len(height)
// 	var total = 0
// 	for i := 0; i < length; i++ {
// 		var maxLeft = height[i]
// 		var maxRight = height[i]
// 		for j := 0; j < i; j++ {
// 			maxLeft = max(maxLeft, height[j])
// 		}
// 		for j := i + 1; j < length; j++ {
// 			maxRight = max(maxRight, height[j])
// 		}
// 		total += min(maxLeft, maxRight) - height[i]
// 	}
// 	return total
// }
