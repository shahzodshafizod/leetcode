package matrices

// https://leetcode.com/problems/max-area-of-island/

// DFS
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down, left

	var dfs func(row int, col int) int

	dfs = func(row int, col int) int {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 0 {
			return 0
		}

		count := 1
		grid[row][col] = 0

		for _, dir := range directions {
			count += dfs(row+dir[0], col+dir[1])
		}

		return count
	}
	maxCount := 0

	for row := range m {
		for col := range n {
			count := dfs(row, col)
			maxCount = max(maxCount, count)
		}
	}

	return maxCount
}

// // BFS
// func maxAreaOfIsland(grid [][]int) int {
// 	var m = len(grid)
// 	var n = len(grid[0])
// 	var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down, left
// 	var maxCount = 0
// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			if grid[row][col] == 0 {
// 				continue
// 			}
// 			var count = 0
// 			var queue = [][2]int{{row, col}}
// 			grid[row][col] = 0
// 			for length := len(queue); length > 0; length = len(queue) {
// 				cell := queue[0]
// 				queue = queue[1:]
// 				count++
// 				for _, dir := range directions {
// 					r := cell[0] + dir[0]
// 					c := cell[1] + dir[1]
// 					if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1 {
// 						queue = append(queue, [2]int{r, c})
// 						grid[r][c] = 0
// 					}
// 				}
// 			}
// 			maxCount = max(maxCount, count)
// 		}
// 	}
// 	return maxCount
// }
