package prefixsums

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/

// Approach #2: Prefix Sum + Binary Search (Ordered Set)
// Time: O(n²m·log(m)) where n = cols, m = rows
// Space: O(m)
func maxSumSubmatrix(matrix [][]int, k int) int {
	// maxSumSubarray finds maximum subarray sum <= k using prefix sum + binary search.
	//
	// For each position i, we want to find the smallest prefix sum
	// that is >= (currentSum - k), so that (currentSum - prefixSum) <= k.
	//
	// We maintain a sorted list of all prefix sums seen so far.
	maxSumSubarray := func(nums []int, k int) int {
		maxSum := math.MinInt
		prefixSums := []int{0} // Sorted list of prefix sums
		currentSum := 0

		for _, num := range nums {
			currentSum += num

			// Find smallest prefix sum >= (currentSum - k)
			// This gives us the subarray sum = currentSum - prefixSum <= k
			target := currentSum - k
			idx := sort.SearchInts(prefixSums, target)

			if idx < len(prefixSums) {
				subarraySum := currentSum - prefixSums[idx]
				if subarraySum > maxSum {
					maxSum = subarraySum
				}
			}

			// Insert current sum in sorted order
			insertIdx := sort.SearchInts(prefixSums, currentSum)
			prefixSums = append(prefixSums, 0)
			copy(prefixSums[insertIdx+1:], prefixSums[insertIdx:])
			prefixSums[insertIdx] = currentSum
		}

		return maxSum
	}

	m, n := len(matrix), len(matrix[0])
	maxSum := math.MinInt

	// Optimize: iterate over smaller dimension
	// If rows >> cols, we iterate over rows instead
	if m > n {
		// Transpose: iterate over rows as column boundaries
		for left := range m {
			rowSums := make([]int, n)
			for right := left; right < m; right++ {
				// Compress columns [left, right] into 1D array
				for col := range n {
					rowSums[col] += matrix[right][col]
				}

				// Find max subarray sum <= k in 1D array
				currentMax := maxSumSubarray(rowSums, k)
				if currentMax > maxSum {
					maxSum = currentMax
				}

				if maxSum == k {
					return k
				}
			}
		}
	} else {
		// Standard: iterate over columns as boundaries
		for left := range n {
			rowSums := make([]int, m)
			for right := left; right < n; right++ {
				// Compress columns [left, right] into 1D array
				for row := range m {
					rowSums[row] += matrix[row][right]
				}

				// Find max subarray sum <= k in 1D array
				currentMax := maxSumSubarray(rowSums, k)
				if currentMax > maxSum {
					maxSum = currentMax
				}

				if maxSum == k {
					return k
				}
			}
		}
	}

	return maxSum
}

// // Approach #1: Brute Force with Prefix Sum
// // Time: O(m²n²)
// // Space: O(mn)
// func maxSumSubmatrix(matrix [][]int, k int) int {
// 	m, n := len(matrix), len(matrix[0])

// 	// Build 2D prefix sum array
// 	// presum[i][j] = sum of rectangle from (0,0) to (i-1,j-1)
// 	presum := make([][]int, m+1)
// 	for i := range presum {
// 		presum[i] = make([]int, n+1)
// 	}

// 	for row := range m {
// 		for col := range n {
// 			presum[row+1][col+1] = matrix[row][col] +
// 				presum[row][col+1] +
// 				presum[row+1][col] -
// 				presum[row][col]
// 		}
// 	}

// 	maxSum := math.MinInt

// 	// Try all possible rectangles
// 	for top := range m {
// 		for left := range n {
// 			for bottom := top; bottom < m; bottom++ {
// 				for right := left; right < n; right++ {
// 					// Calculate sum of rectangle (top,left) to (bottom,right)
// 					rectSum := presum[bottom+1][right+1] -
// 						presum[top][right+1] -
// 						presum[bottom+1][left] +
// 						presum[top][left]

// 					if rectSum <= k && rectSum > maxSum {
// 						maxSum = rectSum
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return maxSum
// }
