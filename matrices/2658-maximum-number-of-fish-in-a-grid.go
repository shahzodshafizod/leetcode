package matrices

// https://leetcode.com/problems/maximum-number-of-fish-in-a-grid/

// Approach: Depth-First Search
// Time: O(mn)
// Space: O(1)
func findMaxFish(grid [][]int) int {
	directions := [5]int{1, 0, -1, 0, 1}
	m, n := len(grid), len(grid[0])

	var dfs func(row int, col int) int

	dfs = func(row int, col int) int {
		fish := grid[row][col]
		grid[row][col] = 0

		var r, c int
		for dir := 1; dir < 5; dir++ {
			r = row + directions[dir-1]
			c = col + directions[dir]

			if min(r, c) < 0 || r == m || c == n || grid[r][c] == 0 {
				continue
			}

			fish += dfs(r, c)
		}

		return fish
	}
	maxfish := 0

	for row := range m {
		for col := range n {
			if grid[row][col] > 0 {
				maxfish = max(maxfish, dfs(row, col))
			}
		}
	}

	return maxfish
}

// // Approach: Depth-First Search (Immutable Input)
// // Time: O(mn)
// // Space: O(mn)
// func findMaxFish(grid [][]int) int {
// 	var directions = [5]int{1, 0, -1, 0, 1}
// 	var m, n = len(grid), len(grid[0])
// 	var seen = make([][]bool, m)
// 	for row := 0; row < m; row++ {
// 		seen[row] = make([]bool, n)
// 	}
// 	var dfs func(row int, col int) int
// 	dfs = func(row int, col int) int {
// 		var fish = grid[row][col]
// 		seen[row][col] = true
// 		var r, c int
// 		for dir := 1; dir < 5; dir++ {
// 			r = row + directions[dir-1]
// 			c = col + directions[dir]
// 			if min(r, c) < 0 || r == m || c == n || grid[r][c] == 0 || seen[r][c] {
// 				continue
// 			}
// 			fish += dfs(r, c)
// 		}
// 		return fish
// 	}
// 	var maxfish = 0
// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			if grid[row][col] > 0 && !seen[row][col] {
// 				maxfish = max(maxfish, dfs(row, col))
// 			}
// 		}
// 	}
// 	return maxfish
// }
