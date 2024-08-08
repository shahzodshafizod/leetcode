package stacks

// https://leetcode.com/problems/largest-rectangle-in-histogram/

// time: O(n)
// space: O(n)
func largestRectangleArea(heights []int) int {
	var n = len(heights)
	var stack = make([][2]int, n) // increasing stack {index, height}
	var top = -1
	var maximal = 0
	var index int
	for idx, height := range heights {
		index = idx
		for top >= 0 && stack[top][1] >= height {
			index = stack[top][0]
			maximal = max(maximal, stack[top][1]*(idx-index))
			top--
		}
		top++
		stack[top][0] = index
		stack[top][1] = height
	}
	var height int
	for top >= 0 {
		index = stack[top][0]
		height = stack[top][1]
		maximal = max(maximal, height*(n-index))
		top--
	}
	return maximal
}

// // time: O(n^2)
// // space: O(1)
// func largestRectangleArea(heights []int) int {
// 	var height, maxheight int
// 	var n = len(heights)
// 	for i := range heights {
// 		height = 10000
// 		for j := i; j < n; j++ {
// 			height = min(height, heights[j])
// 			if height == 0 {
// 				break
// 			}
// 			maxheight = max(maxheight, height*(j-i+1))
// 		}
// 	}
// 	return maxheight
// }
