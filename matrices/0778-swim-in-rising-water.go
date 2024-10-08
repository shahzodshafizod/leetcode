package matrices

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/swim-in-rising-water/

// time: O(n x n)
// space: O(n x n)
func swimInWater(grid [][]int) int {
	var n = len(grid)
	var pq = pkg.NewHeap(
		make([][3]int, 0),
		func(x, y [3]int) bool { return x[0] < y[0] },
	)
	heap.Push(pq, [3]int{grid[0][0], 0, 0})
	grid[0][0] = -1
	var top [3]int
	var row, col, r, c int
	var t = 0
	var directions = [5]int{0, 1, 0, -1, 0}
	for pq.Len() > 0 {
		top = heap.Pop(pq).([3]int)
		t = max(t, top[0])
		row, col = top[1], top[2]
		if row == n-1 && col == n-1 {
			break
		}
		for dir := 0; dir < 4; dir++ {
			r = row + directions[dir]
			c = col + directions[dir+1]
			if min(r, c) >= 0 && max(r, c) < n && grid[r][c] >= 0 {
				heap.Push(pq, [3]int{grid[r][c], r, c})
				grid[r][c] = -1
			}
		}
	}
	return t
}
