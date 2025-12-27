package maths

// https://leetcode.com/problems/range-addition-ii/

// Approach #2: Math (Find intersection of all ranges)
// Time: O(k) - k operations
// Space: O(1) - constant space
func maxCount(m int, n int, ops [][]int) int {
	// If no operations, all cells are 0 (maximum), return total cells
	if len(ops) == 0 {
		return m * n
	}

	// The cells with maximum value are those incremented by ALL operations
	// This is the intersection of all operation ranges: [0, a) x [0, b)
	// The intersection is [0, min(all a)) x [0, min(all b))

	minA := m // Initialize to matrix dimensions
	minB := n

	for _, op := range ops {
		a, b := op[0], op[1]
		if a < minA {
			minA = a
		}

		if b < minB {
			minB = b
		}
	}

	// The number of cells in the intersection
	return minA * minB
}

// // Approach #1: Brute-Force (Simulate the matrix)
// // Time: O(k * m * n) - k operations, each updating up to m*n cells
// // Space: O(m * n) - storing the entire matrix
// func maxCount(m int, n int, ops [][]int) int {
// 	// Create m x n matrix initialized with zeros
// 	matrix := make([][]int, m)
// 	for i := range matrix {
// 		matrix[i] = make([]int, n)
// 	}

// 	// Apply each operation
// 	for _, op := range ops {
// 		a, b := op[0], op[1]
// 		for i := range a {
// 			for j := range b {
// 				matrix[i][j]++
// 			}
// 		}
// 	}

// 	// Find maximum value
// 	maxVal := 0
// 	for i := range m {
// 		for j := range n {
// 			if matrix[i][j] > maxVal {
// 				maxVal = matrix[i][j]
// 			}
// 		}
// 	}

// 	// Count cells with maximum value
// 	count := 0
// 	for i := range m {
// 		for j := range n {
// 			if matrix[i][j] == maxVal {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }
