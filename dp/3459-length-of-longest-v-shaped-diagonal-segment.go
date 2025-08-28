package dp

// https://leetcode.com/problems/length-of-longest-v-shaped-diagonal-segment/

// Approach: Memoization Search
// Time: O(n x m)
// Space: O(n x m)
func lenOfVDiagonal(grid [][]int) int {
	// top-left, top-right, bottom-right, bottom-left
	dirs := [4][2]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	n, m := len(grid), len(grid[0])

	memo := make([][][2][4][3]int, n)
	for idx := range memo {
		memo[idx] = make([][2][4][3]int, m)
	}

	var dfs func(row int, col int, turn int, dir int, target int) int

	dfs = func(row int, col int, turn int, dir int, target int) int {
		if min(row, col) < 0 || row == n || col == m || grid[row][col] != target {
			return 0
		}

		if memo[row][col][turn][dir][target] != 0 {
			return memo[row][col][turn][dir][target]
		}

		nrow, ncol := row+dirs[dir][0], col+dirs[dir][1]

		length := 1 + dfs(nrow, ncol, turn, dir, 2-target)
		if turn == 1 {
			ndir := (dir + 1) % 4
			nrow, ncol = row+dirs[ndir][0], col+dirs[ndir][1]
			length = max(length, 1+dfs(nrow, ncol, 0, ndir, 2-target))
		}

		memo[row][col][turn][dir][target] = length

		return length
	}

	length := 0

	for row := range n {
		for col := range m {
			if grid[row][col] == 1 {
				for dir := range 4 {
					length = max(length, 1+dfs(row+dirs[dir][0], col+dirs[dir][1], 1, dir, 2))
				}
			}
		}
	}

	return length
}
