package matrices

import (
	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-the-safest-path-in-a-grid/

func maximumSafenessFactor(grid [][]int) int {
	var n = len(grid)
	if grid[0][0]+grid[n-1][n-1] > 0 {
		return 0
	}
	var queue = make([][3]int, 0, n*n)
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				queue = append(queue, [3]int{0, row, col})
				grid[row][col] = 0
			} else {
				grid[row][col] = -1
			}
		}
	}
	var inBound = func(row int, col int) bool {
		return min(row, col) >= 0 && max(row, col) < n
	}
	var idx, length = 0, len(queue)
	var directions = [5]int{0, 1, 0, -1, 0}
	var row, col, dist int
	for length-idx > 0 {
		for idx < length {
			for dir := 0; dir < 4; dir++ {
				dist = queue[idx][0]
				row = queue[idx][1] + directions[dir]
				col = queue[idx][2] + directions[dir+1]
				if inBound(row, col) && grid[row][col] == -1 {
					grid[row][col] = dist + 1
					queue = append(queue, [3]int{grid[row][col], row, col})
				}
			}
			idx++
		}
		length = len(queue)
	}
	var newPQ = func() pkg.PQ[[3]int] {
		return pkg.NewPQ(make([][3]int, 0), func(x, y [3]int) bool { return x[0] < y[0] })
	}
	var pq = newPQ()
	pq.Push([3]int{grid[0][0], 0, 0})
	grid[0][0] = -1
	var top [3]int
	for pq.Len() > 0 {
		top = pq.Pop()
		dist = top[0]
		row = top[1]
		col = top[2]
		if row == n-1 && col == n-1 {
			return dist
		}
		for dir := 0; dir < 4; dir++ {
			var r = row + directions[dir]
			var c = col + directions[dir+1]
			if inBound(r, c) && grid[r][c] > 0 {
				grid[r][c] = min(dist, grid[r][c])
				pq.Push([3]int{grid[r][c], r, c})
				grid[r][c] = -1
			}
		}
	}
	return 0
}
