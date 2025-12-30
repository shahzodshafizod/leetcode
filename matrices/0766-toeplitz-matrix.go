package matrices

// https://leetcode.com/problems/toeplitz-matrix/

// Approach: Compare Each Element with Top-Left Neighbor
// A Toeplitz matrix has the property that all diagonals from top-left to bottom-right
// contain the same element. This means for every position (i, j) where i > 0 and j > 0,
// matrix[i][j] should equal matrix[i-1][j-1].
//
// Algorithm:
// 1. Iterate through the matrix starting from row 1, column 1
// 2. For each element, check if it equals its top-left neighbor
// 3. If any element differs from its top-left neighbor, return false
// 4. If all elements match their top-left neighbors, return true
//
// Time: O(m * n) where m is rows, n is columns - visit each element once
// Space: O(1) - only using constant extra space
func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}

	return true
}
