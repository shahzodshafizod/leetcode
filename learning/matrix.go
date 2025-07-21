package learning

import "github.com/shahzodshafizod/leetcode/pkg"

/*
2-D Arrays - Matrices (Matrixes)
BFS (Breadth First Search) and DFS (Depth First Search) In 2D Arrays
*/

type Matrix interface {
	BFS() []int
	DFS() []int
}

type matrix struct {
	grid   [][]int
	height int
	width  int
	dirs   [5]int
}

var _ Matrix = &matrix{}

func NewMatrix(grid [][]int) Matrix {
	return &matrix{
		grid:   grid,
		height: len(grid),
		width:  len(grid[0]),
		dirs:   [5]int{-1, 0, 1, 0, -1},
	}
}

func (m *matrix) BFS() []int {
	seen := make([][]bool, m.height)
	for i := 0; i < m.height; i++ {
		seen[i] = make([]bool, m.width)
	}

	queue := pkg.NewQueue[[2]int]()
	queue.Enqueue([2]int{0, 0})
	seen[0][0] = true

	values := make([]int, 0, m.height*m.width)
	for !queue.Empty() {
		current := queue.Dequeue()
		row, col := current[0], current[1]

		values = append(values, m.grid[row][col])

		for dir := 1; dir < 5; dir++ {
			r := row + m.dirs[dir-1]
			c := col + m.dirs[dir]
			if min(r, c) >= 0 && r < m.height && c < m.width && !seen[r][c] {
				queue.Enqueue([2]int{r, c})
				seen[r][c] = true
			}
		}
	}

	return values
}

func (m *matrix) DFS() []int {
	seen := make([][]bool, m.height)
	for row := 0; row < m.height; row++ {
		seen[row] = make([]bool, m.width)
	}

	var dfs func(row int, col int) []int
	dfs = func(row int, col int) []int {
		if min(row, col) < 0 || row >= m.height || col >= m.width || seen[row][col] {
			return nil
		}
		values := []int{m.grid[row][col]}
		seen[row][col] = true
		for dir := 1; dir < 5; dir++ {
			values = append(values, dfs(row+m.dirs[dir-1], col+m.dirs[dir])...)
		}
		return values
	}

	return dfs(0, 0)
}
