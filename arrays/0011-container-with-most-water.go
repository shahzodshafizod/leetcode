package arrays

/*
Problem:
You are given an array of positive integers where each integer represents the height of a
vertical line on a chart. Find two lines which together with the x-axis forms a container
that would hold the greatest amount of water. Return the area of water it would hold

Step 1: Verify the constraints
	- Does the thickness of the lines affect the area?
		: No, assume they take up no space.
	- Do the left and right sides of the graph count as walls?
		: No, the sides cannot be used to form a container.
	- Does a higher line inside out container affect our area?
		: No, lines inside a container don't affect the area.
	- Can we pick two values if one value is higher in the middle?
		: Yes, the value in the middle won't affect the container.
Step 2: Write out some test cases
	- [7, 1, 2, 3, 9] -> 28 (7,9 -> 7x4)
	- [] -> 0
	- [7] -> 0
	- [6, 9, 3, 4, 5, 8] -> 32 (9,8 -> 8x4)
	- [4, 8, 1, 2, 3, 9}] -> 32 (8,9 -> 8x4)
Step 3: Figure out a solution without code
Step 8: Can we optimize our solution?
*/

// https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	maxArea := 0

	var area int

	for left, right := 0, len(height)-1; left < right; {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}

		maxArea = max(maxArea, area)
	}

	return maxArea
}

// func maxArea(height []int) int {
// 	var maxArea int = 0
// 	var right = len(height) - 1
// 	var left = right - 1
// 	for i := left - 1; i >= 0; i-- {
// 		area1 := min(height[right], height[i]) * (right - i)
// 		area2 := min(height[left], height[i]) * (left - i)
// 		if area1 > maxArea {
// 			maxArea = area1
// 			left = i
// 		}
// 		if area2 > maxArea {
// 			maxArea = area2
// 			right = left
// 			left = i
// 		}
// 	}
// 	return maxArea
// }

// func maxArea(height []int) int {
// 	var maxArea int = 0
// 	var area int
// 	for right := len(height) - 1; right >= 0; right-- {
// 		for left := right - 1; left >= 0; left-- {
// 			area = min(height[right], height[left]) * (right - left)
// 			maxArea = max(maxArea, area)
// 		}
// 	}
// 	return maxArea
// }
