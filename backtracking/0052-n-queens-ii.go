package backtracking

// https://leetcode.com/problems/n-queens-ii/

func totalNQueens(n int) int {
	cols := make([]bool, n)
	diag := make([]bool, 2*n)     // [row+col] as slash /
	backDiag := make([]bool, 2*n) // [n-row+col] as back slash \
	var backtrack func(row int) int
	backtrack = func(row int) int {
		if row == n {
			return 1
		}
		solutions := 0
		for col := 0; col < n; col++ {
			if cols[col] || diag[row+col] || backDiag[n-row+col] {
				continue
			}

			cols[col] = true
			diag[row+col] = true
			backDiag[n-row+col] = true

			solutions += backtrack(row + 1)

			cols[col] = false
			diag[row+col] = false
			backDiag[n-row+col] = false
		}
		return solutions
	}
	return backtrack(0)
}
