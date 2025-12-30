package matrices

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/

// Approach: Staircase/Linear Search (Optimal - Follow-up Solution)
// Key insight: Start from top-right or bottom-left corner
// The matrix is sorted both row-wise and column-wise in non-increasing order
//
// Starting from bottom-left (grid[m-1][0]):
//   - If current element is negative, all elements to the right are also negative
//     (because row is non-increasing), so count them and move up
//   - If current element is non-negative, move right to find negatives
//
// Time: O(m + n) - at most m+n moves (either move up or move right)
// Space: O(1)
func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	count := 0
	row, col := m-1, 0 // Start from bottom-left corner

	for row >= 0 && col < n {
		if grid[row][col] < 0 {
			// All elements from col to n-1 in this row are negative
			count += n - col
			// Move up to check previous row
			row--
		} else {
			// Current element is non-negative, move right
			col++
		}
	}

	return count
}

// // Approach 2: Binary Search on Each Row
// // Since each row is sorted in non-increasing order, use binary search
// // to find the first negative number in each row
// // Time: O(m * log n) where m is rows, n is columns
// // Space: O(1)
// func countNegatives(grid [][]int) int {
// 	n := len(grid[0])
// 	count := 0

// 	for _, row := range grid {
// 		// Binary search to find first negative (first element < 0)
// 		left, right := 0, n-1
// 		firstNeg := n // If no negative found, all are non-negative

// 		for left <= right {
// 			mid := left + (right-left)/2
// 			if row[mid] < 0 {
// 				// Found a negative, but there might be an earlier one
// 				firstNeg = mid
// 				right = mid - 1
// 			} else {
// 				// row[mid] >= 0, search right half
// 				left = mid + 1
// 			}
// 		}

// 		// Count negatives from firstNeg to end of row
// 		count += n - firstNeg
// 	}

// 	return count
// }

// // Approach 1: Brute Force
// // Iterate through every element in the matrix and count negatives
// // Time: O(m * n) where m is rows, n is columns
// // Space: O(1)
// func countNegatives(grid [][]int) int {
// 	count := 0
// 	for _, row := range grid {
// 		for _, num := range row {
// 			if num < 0 {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }
