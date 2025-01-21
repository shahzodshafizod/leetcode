package matrices

// https://leetcode.com/problems/first-completely-painted-row-or-column/

// Approach: Counting
// Time: O(mn)
// Space: O(mn)
func firstCompleteIndex(arr []int, mat [][]int) int {
	var m, n = len(mat), len(mat[0])
	var positions = make([][2]int, m*n+1)
	var row, col int
	for row = 0; row < m; row++ {
		for col = 0; col < n; col++ {
			positions[mat[row][col]] = [2]int{row, col}
		}
	}
	var idx = -1
	var rcount = make([]int, m)
	var ccount = make([]int, n)
	for idx = range arr {
		row, col = positions[arr[idx]][0], positions[arr[idx]][1]
		rcount[row]++
		ccount[col]++
		if rcount[row] == m || ccount[col] == n {
			break
		}
	}
	return idx
}
