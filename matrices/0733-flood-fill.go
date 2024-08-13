package matrices

import "github.com/shahzodshafizod/alkhwarizmi/design"

// https://leetcode.com/problems/flood-fill/

// Approach: Breadth-First Search
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var queue = design.NewQueue[[2]int]()
	var srcColor = image[sr][sc]
	if srcColor != color {
		queue.Enqueue([2]int{sr, sc})
	}
	var m, n = len(image), len(image[0])
	var dirs = [5]int{-1, 0, 1, 0, -1}
	var row, col, r, c int
	for !queue.Empty() {
		var curr = queue.Dequeue()
		row, col = curr[0], curr[1]
		image[row][col] = color
		for d := 1; d < 5; d++ {
			r, c = row+dirs[d-1], col+dirs[d]
			if min(r, c) >= 0 && r < m && c < n && image[r][c] == srcColor {
				queue.Enqueue([2]int{r, c})
			}
		}
	}
	return image
}

// // Approach: Depth-First Search
// func floodFill(image [][]int, sr int, sc int, color int) [][]int {
// 	var m, n = len(image), len(image[0])
// 	var dirs = [5]int{-1, 0, 1, 0, -1}
// 	var dfs func(row int, col int, colorFrom int, colorTo int)
// 	dfs = func(row int, col int, colorFrom int, colorTo int) {
// 		if min(row, col) < 0 || row == m || col == n ||
// 			colorFrom == colorTo || image[row][col] != colorFrom {
// 			return
// 		}
// 		image[row][col] = colorTo
// 		var r, c int
// 		for d := 1; d < 5; d++ {
// 			r, c = row+dirs[d-1], col+dirs[d]
// 			dfs(r, c, colorFrom, colorTo)
// 		}
// 	}
// 	dfs(sr, sc, image[sr][sc], color)
// 	return image
// }
