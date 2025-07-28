package matrices

// https://leetcode.com/problems/convert-1d-array-into-2d-array/

func construct2DArray(original []int, m int, n int) [][]int {
	olen := len(original)
	if m*n != olen {
		return [][]int{}
	}

	matrix := make([][]int, m)
	// for row := 0; row < m; row++ {
	// 	matrix[row] = append([]int{}, original[n*row:n*(row+1)]...)
	// }
	for row, idx := 0, 0; idx < olen; row, idx = row+1, idx+n {
		matrix[row] = append([]int{}, original[idx:idx+n]...)
	}

	return matrix
}
