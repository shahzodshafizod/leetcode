package backtracking

/*
Backtracking
Return ALL solutions
Return a VALID solution amongst all solutions

Problem:
Create a function that solves for any 9x9 (nine by nine) sudoku puzzle given.
Rules:
	- Each of the digits 1-9 must occur exactly once in each row.
	- Each of the digits 1-9 must occur exactly once in each column.
	- Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Step 1: Verify the constraints
	- What happens if the board cannot be solved?
		: Leave the sudoke puzzle as is.
Step 2: Write out some test cases

	5 3 . | . 7 . | . . .
	6 . . | 1 9 5 | . . .
	. 9 8 | . . . | . 6 .
	------+-------+------
	8 . . | . 6 . | . . 3
	4 . . | 8 . 3 | . . 1
	7 . . | . 2 . | . . 6
	------+-------+------
	. 6 . | . . . | 2 8 .
	. . . | 4 1 9 | . . 5
	. . . | . 8 . | . 7 9

	5 3 4 | 6 7 8 | 9 1 2
	6 7 2 | 1 9 5 | 3 4 8
	1 9 8 | 3 4 2 | 5 6 7
	------+-------+------
	8 5 9 | 7 6 1 | 4 2 3
	4 2 6 | 8 5 3 | 7 9 1
	7 1 3 | 9 2 4 | 8 5 6
	------+-------+------
	9 6 1 | 5 3 7 | 2 8 4
	2 8 7 | 4 1 9 | 6 3 5
	3 4 5 | 2 8 6 | 1 7 9

	. . . | . . . | . . .
	. . . | . . . | . 8 .
	. . . | . . . | 1 2 .
	------+-------+------
	. . . | 6 . . | . . .
	. . . | . 9 . | . . .
	. 8 9 | . . . | . 4 .
	------+-------+------
	. . . | . . 6 | . . 2
	. 9 . | . . . | . . .
	. . . | . . . | . . .

	1 2 3 | 4 5 8 | 6 7 9
	4 6 7 | 1 2 9 | 3 8 5
	9 5 8 | 3 6 7 | 1 2 4
	------+-------+------
	2 1 4 | 6 3 5 | 7 9 8
	3 7 5 | 8 9 4 | 2 1 6
	6 8 9 | 2 7 1 | 5 4 3
	------+-------+------
	5 4 1 | 7 8 6 | 9 3 2
	7 9 2 | 5 4 3 | 8 6 1
	8 3 6 | 9 1 2 | 4 5 7

	. . . | . . . | . . .
	. . . | . . . | . . .
	. . . | . . . | . . .
	------+-------+------
	. . . | . . . | . . .
	. . . | . . . | . . .
	. . . | . . . | . . .
	------+-------+------
	. . . | . . . | . . .
	. . . | . . . | . . .
	. . . | . . . | . . .

	1 2 3 | 4 5 6 | 7 8 9
	4 5 6 | 7 8 9 | 1 2 3
	7 8 9 | 1 2 3 | 4 5 6
	------+-------+------
	2 1 4 | 3 6 5 | 8 9 7
	3 6 5 | 8 9 7 | 2 1 4
	8 9 7 | 2 1 4 | 3 6 5
	------+-------+------
	5 3 1 | 6 4 2 | 9 7 8
	6 4 2 | 9 7 8 | 5 3 1
	9 7 8 | 5 3 1 | 6 4 2
*/

// https://leetcode.com/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}

			if !sudokuIsValid(board, row, col, board[row][col]) {
				return
			}
		}
	}

	sudokuBacktrack(board, 0, 0)
}

func sudokuIsValid(board [][]byte, row, col int, num byte) bool {
	blkrow, blkcol := (row/3)*3, (col/3)*3

	for i := 0; i < 9; i++ {
		if i != row && board[i][col] == num {
			return false
		}

		if i != col && board[row][i] == num {
			return false
		}

		r := blkrow + i/3
		c := blkcol + i%3

		if r != row && c != col && board[r][c] == num {
			return false
		}
	}

	return true
}

func sudokuBacktrack(board [][]byte, row, col int) bool {
	if row == 8 && col == 9 {
		return true
	}

	if col == 9 {
		row++
		col = 0
	}

	if board[row][col] != '.' {
		return sudokuBacktrack(board, row, col+1)
	}

	var num byte
	for num = '1'; num <= '9'; num++ {
		if sudokuIsValid(board, row, col, num) {
			board[row][col] = num

			if sudokuBacktrack(board, row, col+1) {
				return true
			}

			board[row][col] = '.'
		}
	}

	return false // unreachable if solvable
}

// func solveSudoku(board [][]byte) {
// 	var (
// 		rows  = [9]map[byte]bool{}
// 		cols  = [9]map[byte]bool{}
// 		boxes = [9]map[byte]bool{}
// 	)
// 	for i := 0; i < 9; i++ {
// 		rows[i] = make(map[byte]bool)
// 		cols[i] = make(map[byte]bool)
// 		boxes[i] = make(map[byte]bool)
// 	}

// 	var value byte
// 	for row := 0; row < 9; row++ {
// 		for col := 0; col < 9; col++ {
// 			value = board[row][col]
// 			if value == '.' {
// 				continue
// 			}
// 			var boxID = 3*(row/3) + col/3
// 			if !sudokuIsValid(rows[row], cols[col], boxes[boxID], value) {
// 				return
// 			}
// 			rows[row][value] = true
// 			cols[col][value] = true
// 			boxes[boxID][value] = true
// 		}
// 	}

// 	sudokuBacktrack(board, rows, cols, boxes, 0, 0)
// }

// func sudokuIsValid(row, col, box map[byte]bool, num byte) bool {
// 	return !row[num] && !col[num] && !box[num]
// }

// func sudokuBacktrack(board [][]byte, rows, cols, boxes [9]map[byte]bool, row, col int) bool {
// 	if row == 8 && col == 9 {
// 		return true
// 	}
// 	if col == 9 {
// 		row++
// 		col = 0
// 	}

// 	if board[row][col] != '.' {
// 		return sudokuBacktrack(board, rows, cols, boxes, row, col+1)
// 	}

// 	var boxID = 3*(row/3) + col/3
// 	var num byte
// 	for num = '1'; num <= '9'; num++ {
// 		if sudokuIsValid(rows[row], cols[col], boxes[boxID], num) {
// 			board[row][col] = num
// 			rows[row][num] = true
// 			cols[col][num] = true
// 			boxes[boxID][num] = true

// 			if sudokuBacktrack(board, rows, cols, boxes, row, col+1) {
// 				return true
// 			}

// 			board[row][col] = '.'
// 			rows[row][num] = false
// 			cols[col][num] = false
// 			boxes[boxID][num] = false
// 		}
// 	}
// 	return false // unreachable if solvable
// }
