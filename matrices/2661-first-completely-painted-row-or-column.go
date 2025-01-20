package matrices

// https://leetcode.com/problems/first-completely-painted-row-or-column/

func firstCompleteIndex(arr []int, mat [][]int) int {
	var m, n = len(mat), len(mat[0])
	var positions = make([]int, m*n+1)
	for idx, num := range arr {
		positions[num] = idx
	}
	var idx = m * n
	var maxIdx int
	for row := 0; row < m; row++ {
		maxIdx = 0
		for col := 0; col < n; col++ {
			maxIdx = max(maxIdx, positions[mat[row][col]])
		}
		idx = min(idx, maxIdx)
	}
	for col := 0; col < n; col++ {
		maxIdx = 0
		for row := 0; row < m; row++ {
			maxIdx = max(maxIdx, positions[mat[row][col]])
		}
		idx = min(idx, maxIdx)
	}
	return idx
}
