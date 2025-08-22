package matrices

// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/

func minimumArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	left, right := n, 0

	top, down := m, 0
	for row := range m {
		for col := range n {
			if grid[row][col] == 1 {
				left = min(left, col)
				right = max(right, col)
				top = min(top, row)
				down = max(down, row)
			}
		}
	}

	if left == n {
		return 0
	}

	height := down - top + 1
	width := right - left + 1

	return height * width
}
