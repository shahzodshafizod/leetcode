package matrices

// https://leetcode.com/problems/reshape-the-matrix/

// Approach: Flatten and Rebuild
// Key insight: Treat 2D matrix as 1D array by mapping indices.
// For position (i,j) in m×n matrix: index = i*n + j
// For index k in r×c matrix: row = k/c, col = k%c
//
// Strategy:
// 1. Check if reshape is valid: m*n must equal r*c
// 2. Iterate through original matrix in row-major order
// 3. Place each element in new matrix using index mapping
//
// Time Complexity: O(m*n) where m,n are dimensions of original matrix
// Space Complexity: O(r*c) for the result matrix
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	m, n := len(mat), len(mat[0])

	// Check if reshape is possible
	if m*n != r*c {
		return mat
	}

	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	// Traverse original matrix and fill result
	for i := range m * n {
		// Map 1D index to original 2D position
		origRow := i / n
		origCol := i % n

		// Map 1D index to new 2D position
		newRow := i / c
		newCol := i % c

		result[newRow][newCol] = mat[origRow][origCol]
	}

	return result
}
