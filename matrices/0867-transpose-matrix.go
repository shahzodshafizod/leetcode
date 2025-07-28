package matrices

// https://leetcode.com/problems/transpose-matrix/

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return [][]int{}
	}

	n := len(matrix[0])
	if n == 0 {
		return [][]int{}
	}

	if m == n { // in-place transpose
		for index := 1; index < m; index++ {
			i, j := 0, index
			for ; i < j; i, j = i+1, j-1 {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}

		for index := m - 2; index > 0; index-- {
			i, j := index, m-1
			for ; i < j; i, j = i+1, j-1 {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}

		return matrix
	}

	transposed := make([][]int, n)
	for col := 0; col < n; col++ {
		transposed[col] = make([]int, m)
		for row := 0; row < m; row++ {
			transposed[col][row] = matrix[row][col]
		}
	}

	return transposed
}
