package graphs

// https://leetcode.com/problems/making-a-large-island/

// Approach: Depth-First Search
// Time: O(nn)
// Space: O(nn)
func largestIsland(grid [][]int) int {
	n := len(grid)
	ids := make([][]int, n)

	var row, col int

	for row = range n {
		ids[row] = make([]int, n)
		for col = range n {
			ids[row][col] = -1
		}
	}

	directions := [5]int{1, 0, -1, 0, 1}

	var dfs func(row int, col int, id int) int

	dfs = func(row int, col int, id int) int {
		ids[row][col] = id
		area := 1

		for dir := 1; dir < 5; dir++ {
			r := row + directions[dir-1]
			c := col + directions[dir]

			if min(r, c) >= 0 && max(r, c) < n &&
				grid[r][c] == 1 && ids[r][c] == -1 {
				area += dfs(r, c, id)
			}
		}

		return area
	}

	var areas []int

	var zeroes [][2]int

	id := 0
	maxArea := 0

	for row = range n {
		for col = range n {
			if grid[row][col] == 0 {
				zeroes = append(zeroes, [2]int{row, col})
			} else if ids[row][col] == -1 {
				areas = append(areas, dfs(row, col, id))
				maxArea = max(maxArea, areas[id])
				id++
			}
		}
	}

	if len(areas) == 0 {
		return 1
	}

	for idx := range zeroes {
		row, col = zeroes[idx][0], zeroes[idx][1]
		neighbors := make(map[int]struct{})

		for dir := 1; dir < 5; dir++ {
			r := row + directions[dir-1]
			c := col + directions[dir]

			if min(r, c) >= 0 && max(r, c) < n && grid[r][c] == 1 {
				neighbors[ids[r][c]] = struct{}{}
			}
		}

		area := 1
		for nid := range neighbors {
			area += areas[nid]
		}

		maxArea = max(maxArea, area)
	}

	return maxArea
}
