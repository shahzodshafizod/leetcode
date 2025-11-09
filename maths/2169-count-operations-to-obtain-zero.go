package maths

// https://leetcode.com/problems/count-operations-to-obtain-zero/

// Approach #2: Euclidean Algorithm
// Time: O(log(max(num1, num2)))
// Space: O(1)
func countOperations(num1 int, num2 int) int {
	var ops int
	for num1 != 0 && num2 != 0 {
		ops += num1 / num2
		num1 %= num2
		num1, num2 = num2, num1
	}

	return ops
}

// // Approach #1: Simulation
// // Time: O(max(num1, num2))
// // Space: O(1)
// func countOperations(num1 int, num2 int) int {
// 	var ops int

// 	for num1 != 0 && num2 != 0 {
// 		if num1 >= num2 {
// 			num1 -= num2
// 		} else {
// 			num2 -= num1
// 		}

// 		ops++
// 	}

// 	return ops
// }
