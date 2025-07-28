package matrices

/*
Modify the grid such that it either:
	- Contains no islands, or
	- Contains more than one island.
*/

// https://leetcode.com/problems/minimum-number-of-days-to-disconnect-island/

// Approach: Brute Force
// Time: O((MxN)^2)
// Space: O(MxN)
func minDays(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [5]int{-1, 0, 1, 0, -1}

	var dfs func(row int, col int, visited [][]bool) int

	dfs = func(row int, col int, visited [][]bool) int {
		if min(row, col) < 0 || row == m || col == n ||
			grid[row][col] == 0 || visited[row][col] {
			return 0
		}

		visited[row][col] = true

		var r, c int
		for d := 1; d < 5; d++ {
			r, c = row+dirs[d-1], col+dirs[d]
			dfs(r, c, visited)
		}

		return 1
	}

	countIslands := func() int {
		visited := make([][]bool, m)
		for idx := range visited {
			visited[idx] = make([]bool, n)
		}

		count := 0

		for row := 0; row < m; row++ {
			for col := 0; col < n; col++ {
				count += dfs(row, col, visited)
			}
		}

		return count
	}
	if countIslands() != 1 {
		return 0
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == 0 {
				continue
			}

			grid[row][col] = 0

			if countIslands() != 1 {
				return 1
			}

			grid[row][col] = 1
		}
	}

	return 2
}

// Think about Tarjan's Algorithm
// Its time & space complexities are O(M x N)
