package backtracking

// https://leetcode.com/problems/path-with-maximum-gold/

// time: O(m x n x 4^(m x n))
// space: O(4^(m x n))
// OR 3^(m x n) since the path never goes back from the direction it came from
func getMaximumGold(grid [][]int) int {
	// east[0, 1], south[1, 0], west[0, -1], north[-1, 0]
	directions := [5]int{0, 1, 0, -1, 0}
	m, n := len(grid), len(grid[0])

	var dfs func(row int, col int) int

	dfs = func(row int, col int) int {
		gold := grid[row][col]
		grid[row][col] = 0
		maxgold := 0

		for idx := 0; idx < 4; idx++ {
			r, c := row+directions[idx], col+directions[idx+1]
			if min(r, c) >= 0 && r < m && c < n && grid[r][c] != 0 {
				maxgold = max(maxgold, dfs(r, c))
			}
		}

		grid[row][col] = gold

		return maxgold + gold
	}
	maxgold := 0

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] != 0 {
				maxgold = max(maxgold, dfs(row, col))
			}
		}
	}

	return maxgold
}
