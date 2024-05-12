package matrices

// https://leetcode.com/problems/largest-local-values-in-a-matrix/

func largestLocal(grid [][]int) [][]int {
	var n = len(grid)
	var maxLocal = make([][]int, n-2)
	for r := 0; r < n-2; r++ {
		maxLocal[r] = make([]int, n-2)
		for c := 0; c < n-2; c++ {
			maxLocal[r][c] = max(
				max(grid[r][c], grid[r][c+1], grid[r][c+2]),
				max(grid[r+1][c], grid[r+1][c+1], grid[r+1][c+2]),
				max(grid[r+2][c], grid[r+2][c+1], grid[r+2][c+2]),
			)
		}
	}
	return maxLocal
}
