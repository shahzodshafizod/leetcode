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
	// Check if it's valid and set existed values
	var (
		rows, cols, boxes [9]int
		bit, boxID        int
	)

	for row := range 9 {
		for col := range 9 {
			if board[row][col] != '.' {
				bit = 1 << int(board[row][col]-'1')
				boxID = (row/3)*3 + col/3
				rows[row] ^= bit
				cols[col] ^= bit
				boxes[boxID] ^= bit
			}
		}
	}

	// Solve
	var backtrack func(row int, col int) bool

	backtrack = func(row int, col int) bool {
		if row == 9 {
			return true
		}

		col++
		if col == 9 {
			return backtrack(row+1, -1)
		}

		if board[row][col] != '.' {
			return backtrack(row, col)
		}

		boxID := row/3*3 + col/3

		var bit int
		for num := range 9 {
			bit = 1 << num
			if rows[row]&bit != 0 || cols[col]&bit != 0 || boxes[boxID]&bit != 0 {
				continue
			}

			rows[row] ^= bit
			cols[col] ^= bit
			boxes[boxID] ^= bit

			board[row][col] = byte('1' + num)
			if backtrack(row, col) {
				return true
			}

			rows[row] ^= bit
			cols[col] ^= bit
			boxes[boxID] ^= bit
			board[row][col] = '.'
		}

		return false
	}

	backtrack(0, -1)
}
