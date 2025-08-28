package matrices

// https://leetcode.com/problems/diagonal-traverse/

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])

	diag := make([][]int, m+n)
	for row := range m {
		for col := range n {
			diag[row+col] = append(diag[row+col], mat[row][col])
		}
	}

	result := make([]int, 0, m*n)

	for idx, line := range diag {
		// each line is reversed, so we reverse again
		if idx&1 == 0 {
			for k := len(line) - 1; k >= 0; k-- {
				result = append(result, line[k])
			}
		} else {
			result = append(result, line...)
		}
	}

	return result
}
