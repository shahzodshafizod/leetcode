package greedy

// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/

// Approach #2: Greedy
// Time: O(n)
// Space: O(1)
func minCost1578(colors string, neededTime []int) int {
	var res, maxCost int

	for i := range colors {
		if i > 0 && colors[i-1] != colors[i] {
			maxCost = 0
		}

		res += min(maxCost, neededTime[i])
		maxCost = max(maxCost, neededTime[i])
	}

	return res
}

// // Approach #1: Stack
// // Time: O(n)
// // Space: O(n)
// func minCost1578(colors string, neededTime []int) int {
// 	n := len(colors)
// 	stack := make([]int, n)
// 	size := 0
// 	time := 0

// 	for i := range n {
// 		if size > 0 && colors[stack[size-1]] == colors[i] {
// 			if neededTime[stack[size-1]] < neededTime[i] {
// 				time += neededTime[stack[size-1]]
// 				stack[size-1] = i
// 			} else {
// 				time += neededTime[i]
// 			}
// 		} else {
// 			stack[size] = i
// 			size++
// 		}
// 	}

// 	return time
// }
