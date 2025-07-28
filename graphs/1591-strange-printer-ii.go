package graphs

// https://leetcode.com/problems/strange-printer-ii/

// Approach: Topological Sort
// Time: O(mn), m=# of ROWS, n=# of COLS
// Space: O(c), c=# of colors
func isPrintable(targetGrid [][]int) bool {
	border := make(map[int][4]int)

	for row := range targetGrid { // O(m)
		for col, color := range targetGrid[row] { // O(n)
			if _, found := border[color]; !found {
				border[color] = [4]int{row, row, col, col}
			} else {
				border[color] = [4]int{
					border[color][0],
					max(border[color][1], row),
					min(border[color][2], col),
					max(border[color][3], col),
				}
			}
		}
	}

	seen := make(map[int]bool)

	var dfs func(row int, col int) bool // O(mn)

	dfs = func(row int, col int) bool {
		color := targetGrid[row][col]
		if color == 0 {
			return true
		}

		if seen[color] {
			return false
		}

		seen[color] = true

		for r := border[color][0]; r <= border[color][1]; r++ {
			for c := border[color][2]; c <= border[color][3]; c++ {
				if targetGrid[r][c] != color && !dfs(r, c) {
					return false
				}

				targetGrid[r][c] = 0
			}
		}

		return true
	}

	for row := range targetGrid {
		for col := range targetGrid[row] {
			if !dfs(row, col) {
				return false
			}
		}
	}

	return true
}
