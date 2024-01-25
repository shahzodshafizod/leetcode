package matrices

/*
2-D Arrays - Matrices (Matrixes)
BFS (Breadth First Search) and DFS (Depth First Search) In 2D Arrays
*/

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down and left

func traversalDFS(matrix [][]int) []int {
	var height = len(matrix)
	if height == 0 {
		return nil
	}
	var width = len(matrix[0])
	if width == 0 {
		return nil
	}

	var seen = make([][]bool, height)
	for row := 0; row < height; row++ {
		seen[row] = make([]bool, width)
	}

	return dfs(matrix, seen, height, width, 0, 0)
}

func dfs(matrix [][]int, seen [][]bool, height int, width int, row int, col int) []int {
	if row < 0 || col < 0 || row >= height || col >= width || seen[row][col] {
		return []int{}
	}
	var values = []int{matrix[row][col]}
	seen[row][col] = true
	for _, direction := range directions {
		values = append(values, dfs(matrix, seen, height, width, row+direction[0], col+direction[1])...)
	}
	return values
}

func traversalBFS(matrix [][]int) []int {
	var height = len(matrix)
	if height == 0 {
		return nil
	}
	var width = len(matrix[0])
	if width == 0 {
		return nil
	}

	var seen = make([][]bool, height)
	for i := 0; i < height; i++ {
		seen[i] = make([]bool, width)
	}

	var values = make([]int, height*width)
	var index = 0

	var queue = [][2]int{{0, 0}}
	seen[0][0] = true
	for len(queue) > 0 {
		var row, col = queue[0][0], queue[0][1]

		values[index] = matrix[row][col]
		index++

		for _, direction := range directions {
			var r, c = row + direction[0], col + direction[1]
			if r >= 0 && c >= 0 && r < height && c < width && !seen[r][c] {
				queue = append(queue, [2]int{r, c})
				seen[r][c] = true
			}
		}

		queue = queue[1:]
	}

	return values
}
