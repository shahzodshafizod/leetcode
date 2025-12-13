package maths

import "math"

// https://leetcode.com/problems/arranging-coins/

// Approach #3: Mathematical Formula (Quadratic Equation)
// Time: O(1) - constant time calculation
// Space: O(1) - constant space
func arrangeCoins(n int) int {
	// We need k such that 1 + 2 + ... + k <= n
	// Using formula: k * (k + 1) / 2 <= n
	// Solving quadratic: k^2 + k - 2n <= 0
	// Using quadratic formula: k = (-1 + sqrt(1 + 8n)) / 2
	return int((-1 + math.Sqrt(1+8*float64(n))) / 2)
}

// // Approach #2: Binary Search
// // Time: O(log n) - binary search
// // Space: O(1) - constant space
// func arrangeCoins(n int) int {
// 	// For each mid, check if mid * (mid + 1) / 2 <= n
// 	left, right := 0, n
// 	result := 0

// 	for left <= right {
// 		mid := (left + right) / 2
// 		total := mid * (mid + 1) / 2

// 		if total <= n {
// 			result = mid
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	return result
// }

// // Approach #1: Brute Force - Simulation
// // Time: O(sqrt(n)) - approximately sqrt(n) rows
// // Space: O(1) - constant space
// func arrangeCoins(n int) int {
// 	row := 0
// 	for row+1 <= n {
// 		row++
// 		n -= row
// 	}

// 	return row
// }
