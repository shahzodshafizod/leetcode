package matrices

// https://leetcode.com/problems/set-matrix-zeroes/

func setZeroes(matrix [][]int) {
	var m, n = len(matrix), len(matrix[0])
	var firstRowZero, firstColZero = false, false
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				if row == 0 {
					firstRowZero = true
				} else {
					matrix[row][0] = 0
				}

				if col == 0 {
					firstColZero = true
				} else {
					matrix[0][col] = 0
				}
			}
		}
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
	}
	if firstRowZero {
		for col := 0; col < n; col++ {
			matrix[0][col] = 0
		}
	}
	if firstColZero {
		for row := 0; row < m; row++ {
			matrix[row][0] = 0
		}
	}
}

// func setZeroes(matrix [][]int) {
// 	var m = len(matrix)
// 	if m == 0 {
// 		return
// 	}
// 	var n = len(matrix[0])
// 	if n == 0 {
// 		return
// 	}

// 	// -2^31 <= matrix[i][j] <= 2^31 - 1
// 	// -2147483648 <= matrix[i][j] <= 2147483647
// 	// so we'll first set to 2^31, and then 2^31 to 0
// 	// 2^31 = 2147483648

// 	var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down and left

// 	const TEMPVALUE = 2147483648
// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			if matrix[row][col] == 0 {
// 				for _, direction := range directions {
// 					r, c := row, col
// 					for r >= 0 && c >= 0 && r < m && c < n {
// 						if matrix[r][c] != 0 {
// 							matrix[r][c] = TEMPVALUE
// 						}
// 						r += direction[0]
// 						c += direction[1]
// 					}
// 				}
// 			}
// 		}
// 	}

// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			if matrix[row][col] == TEMPVALUE {
// 				matrix[row][col] = 0
// 			}
// 		}
// 	}
// }

// func setZeroes(matrix [][]int) {
// 	var m = len(matrix)
// 	if m == 0 {
// 		return
// 	}
// 	var n = len(matrix[0])
// 	if n == 0 {
// 		return
// 	}

// 	var rows = make(map[int]bool)
// 	var cols = make(map[int]bool)
// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			if matrix[row][col] == 0 {
// 				rows[row] = true
// 				cols[col] = true
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
