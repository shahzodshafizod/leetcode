package matrices

/*
rowSum = [5,7,10]
colSum = [8,6,8]

          [ 8][6][8]    [8][ 6][8]    [8][6][8]
matrix = [  |                |               |
[ 5]      [ 5][0][0]    [5][ 0][0]    [5][0][0]
[ 7]      [ 7][0][0] => [3][ 4][0] => [3][4][0]
[10]      [10][0][0]    [0][10][0]    [0][2][8]
         ]
*/

// https://leetcode.com/problems/find-valid-matrix-given-row-and-column-sums/

// Approach: Doesn't change the input arrays
// time: O(M x N)
// space: O(1)
func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	matrix := make([][]int, m)
	for row := 0; row < m; row++ {
		matrix[row] = make([]int, n)
		matrix[row][0] = rowSum[row]
	}
	var sum int
	for col := 1; col < n; col++ {
		sum = 0
		for row := 0; row < m; row++ {
			sum += matrix[row][col-1]
			if sum > colSum[col-1] {
				matrix[row][col] = sum - colSum[col-1]
				matrix[row][col-1] -= matrix[row][col]
				sum -= matrix[row][col]
			}
		}
	}
	return matrix
}

// // Approach: Changes the input arrays
// // time: O(M x N)
// // space: O(1)
// func restoreMatrix(rowSum []int, colSum []int) [][]int {
// 	var m, n = len(rowSum), len(colSum)
// 	var matrix = make([][]int, m)
// 	for idx := range matrix {
// 		matrix[idx] = make([]int, n)
// 	}
// 	var row, col = 0, 0
// 	for row < m && col < n {
// 		matrix[row][col] = min(rowSum[row], colSum[col])
// 		rowSum[row] -= matrix[row][col]
// 		colSum[col] -= matrix[row][col]
// 		if rowSum[row] == 0 {
// 			row++
// 		} else {
// 			col++
// 		}
// 	}
// 	return matrix
// }
