package maths

import "sort"

// https://leetcode.com/problems/maximum-product-of-three-numbers/

// Approach #2: Sorting + Math
// Time: O(n log n) - sorting dominates
// Space: O(1) - constant space (ignoring sort space)
func maximumProduct(nums []int) int {
	// Sort the array
	sort.Ints(nums)
	n := len(nums)

	// Maximum product can be either:
	// 1. Product of three largest numbers (all positive or all negative)
	// 2. Product of two smallest (most negative) and largest (if negatives exist)

	// Case 1: Three largest numbers
	product1 := nums[n-1] * nums[n-2] * nums[n-3]

	// Case 2: Two smallest (negative) * largest (positive)
	product2 := nums[0] * nums[1] * nums[n-1]

	if product1 > product2 {
		return product1
	}

	return product2
}

// // Approach #1: Brute-Force (Try all combinations)
// // Time: O(n^3) - three nested loops
// // Space: O(1) - constant space
// func maximumProduct(nums []int) int {
// 	n := len(nums)
// 	maxProduct := -1 << 31

// 	// Try all possible triplets
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			for k := j + 1; k < n; k++ {
// 				product := nums[i] * nums[j] * nums[k]
// 				if product > maxProduct {
// 					maxProduct = product
// 				}
// 			}
// 		}
// 	}

// 	return maxProduct
// }
