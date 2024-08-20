package backtracking

// https://leetcode.com/problems/n-queens/

func solveNQueens(n int) [][]string {
	var solutions = make([][]string, 0)
	var board = make([][]byte, n)
	for row := range board {
		board[row] = make([]byte, n)
		for col := range board[row] {
			board[row][col] = '.'
		}
	}
	var cols = make([]bool, n)      // (row)
	var negDiag = make([]bool, 2*n) // (row - col)
	var posDiag = make([]bool, 2*n) // (row + col)
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var solution = make([]string, n)
			for idx := range solution {
				solution[idx] = string(board[idx])
			}
			solutions = append(solutions, solution)
			return
		}
		for col := 0; col < n; col++ {
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
