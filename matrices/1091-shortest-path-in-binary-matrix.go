package matrices

// https://leetcode.com/problems/shortest-path-in-binary-matrix/

// BFS (DFS works incorrectly)
func shortestPathBinaryMatrix(grid [][]int) int {
	directions := [8][2]int{{-1, 0}, {0, -1}, {-1, 1}, {1, -1}, {0, 1}, {1, 0}, {1, 1}, {-1, -1}}
	m := len(grid)
	n := len(grid[0])
	queue := make([][2]int, 0)
	if grid[0][0] == 0 {
		queue = append(queue, [2]int{0, 0})
		grid[0][0] = 1
	}
	var cell [2]int
	var row, col, r, c int
	count := 0
	for length := len(queue); length > 0; length = len(queue) {
		count++
		for i := 0; i < length; i++ {
			cell = queue[i]
			row = cell[0]
			col = cell[1]
			if row == m-1 && col == n-1 {
				return count
			}
			for _, dir := range directions {
				r = row + dir[0]
				c = col + dir[1]
				if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 0 {
					queue = append(queue, [2]int{r, c})
					grid[r][c] = 1
				}
			}
		}
		queue = queue[length:]
	}
	return -1
}
