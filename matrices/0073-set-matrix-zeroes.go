package matrices

// https://leetcode.com/problems/set-matrix-zeroes/

// Approach #3
// Time: O(mn)
// Space: O(1)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstCol := false

	for row := range m {
		if matrix[row][0] == 0 {
			firstCol = true
		}

		for col := 1; col < n; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	for row := m - 1; row >= 0; row-- {
		for col := n - 1; col >= 1; col-- {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}

		if firstCol {
			matrix[row][0] = 0
		}
	}
}

// // Approach #2
// // Time: O(mn)
// // Space: O(n)
// func setZeroes(matrix [][]int) {
// 	var m = len(matrix)
// 	var n = len(matrix[0])
// 	var cols = make(map[int]struct{})
// 	var zeroRow bool
// 	for row := 0; row < m; row++ {
// 		zeroRow = false
// 		for col := 0; col < n; col++ {
// 			if matrix[row][col] == 0 {
// 				cols[col] = struct{}{}
// 				zeroRow = true
// 			}
// 		}
// 		if zeroRow {
// 			for col := 0; col < n; col++ {
// 				matrix[row][col] = 0
// 			}
// 		}
// 	}
// 	for col := range cols {
// 		for row := 0; row < m; row++ {
// 			matrix[row][col] = 0
// 		}
// 	}
// }

// // Approach #1
// // Time: O(mn)
// // Space: O(m+n)
// func setZeroes(matrix [][]int) {
// 	var m = len(matrix)
// 	var n = len(matrix[0])
// 	var rows = make(map[int]struct{})
// 	var cols = make(map[int]struct{})
// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			if matrix[row][col] == 0 {
// 				rows[row] = struct{}{}
// 				cols[col] = struct{}{}
// 			}
// 		}
// 	}
// 	for row := range rows {
// 		for col := 0; col < n; col++ {
// 			matrix[row][col] = 0
// 		}
// 	}
// 	for col := range cols {
// 		for row := 0; row < m; row++ {
// 			matrix[row][col] = 0
// 		}
// 	}
// }
