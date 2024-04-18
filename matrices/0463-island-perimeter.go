package matrices

// https://leetcode.com/problems/island-perimeter/

func islandPerimeter(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	var blocks, borders = 0, 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == 1 {
				blocks += 4
				if row < m-1 && grid[row+1][col] == 1 {
					// count down/up neighbors
					borders += 2
				}
				if col < n-1 && grid[row][col+1] == 1 {
					// count right/left borders
					borders += 2
				}
			}
		}
	}
	return blocks - borders
}

// func islandPerimeter(grid [][]int) int {
// 	var m, n = len(grid), len(grid[0])
// 	var dfs func(row int, col int) int
// 	dfs = func(row int, col int) int {
// 		if row < 0 || row == m || col < 0 || col == n || grid[row][col] == 0 {
// 			return 1 // counting borders with water
// 		}
// 		if grid[row][col] == 2 {
// 			return 0
// 		}
// 		grid[row][col] = 2
// 		var perimeter = 0
// 		for _, direction := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
// 			perimeter += dfs(row+direction[0], col+direction[1])
// 		}
// 		return perimeter
// 	}
// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			if grid[row][col] == 1 {
// 				return dfs(row, col)
// 			}
// 		}
// 	}
// 	return 0
// }
