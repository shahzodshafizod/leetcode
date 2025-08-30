package bits

// https://leetcode.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	var (
		rows, cols, boxes [9]int
		bit, idx          int
	)

	for row := range 9 {
		for col := range 9 {
			if board[row][col] != '.' {
				idx = row/3*3 + col/3

				bit = 1 << int(board[row][col]-'0')
				if rows[row]&bit != 0 || cols[col]&bit != 0 || boxes[idx]&bit != 0 {
					return false
				}

				rows[row] |= bit
				cols[col] |= bit
				boxes[idx] |= bit
			}
		}
	}

	return true
}
