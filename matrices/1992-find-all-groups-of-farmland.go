package matrices

// https://leetcode.com/problems/find-all-groups-of-farmland/

func findFarmland(land [][]int) [][]int {
	var m, n = len(land), len(land[0])
	var dfs func(row int, col int) (int, int)
	dfs = func(row int, col int) (int, int) {
		land[row][col] = 0
		var r, c int
		for _, next := range [2][2]int{
			{row + 1, col}, // down
			{row, col + 1}, // right
		} {
			r, c = next[0], next[1]
			if r >= 0 && r < m && c >= 0 && c < n && land[r][c] == 1 {
				r, c = dfs(r, c)
				row = max(row, r)
				col = max(col, c)
			}
		}
		return row, col
	}
	var farmland = make([][]int, 0)
	var r2, c2 int
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if land[row][col] == 1 {
				r2, c2 = dfs(row, col)
				farmland = append(farmland, []int{row, col, r2, c2})
				col = c2
			}
		}
	}
	return farmland
}
