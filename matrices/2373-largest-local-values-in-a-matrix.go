package matrices

// https://leetcode.com/problems/largest-local-values-in-a-matrix/

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	maxLocal := make([][]int, n-2)

	for row := 0; row < n-2; row++ {
		maxLocal[row] = make([]int, n-2)
		for col := 0; col < n-2; col++ {
			maxLocal[row][col] = max(
				grid[row][col], grid[row][col+1], grid[row][col+2],
				grid[row+1][col], grid[row+1][col+1], grid[row+1][col+2],
				grid[row+2][col], grid[row+2][col+1], grid[row+2][col+2],
			)
		}
	}

	return maxLocal
}
