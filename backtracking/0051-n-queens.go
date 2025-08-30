package backtracking

// https://leetcode.com/problems/n-queens/

func solveNQueens(n int) [][]string {
	solutions := make([][]string, 0)

	board := make([][]byte, n)
	for row := range board {
		board[row] = make([]byte, n)
		for col := range board[row] {
			board[row][col] = '.'
		}
	}

	cols := make([]bool, n)      // (row)
	negDiag := make([]bool, 2*n) // (row - col)
	posDiag := make([]bool, 2*n) // (row + col)

	var backtrack func(row int)

	backtrack = func(row int) {
		if row == n {
			solution := make([]string, n)
			for idx := range solution {
				solution[idx] = string(board[idx])
			}

			solutions = append(solutions, solution)

			return
		}

		for col := range n {
			if cols[col] || negDiag[n-row+col] || posDiag[row+col] {
				continue
			}

			board[row][col] = 'Q'
			cols[col] = true
			negDiag[n-row+col] = true
			posDiag[row+col] = true

			backtrack(row + 1)

			board[row][col] = '.'
			cols[col] = false
			negDiag[n-row+col] = false
			posDiag[row+col] = false
		}
	}
	backtrack(0)

	return solutions
}
