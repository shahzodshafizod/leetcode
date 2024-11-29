package graphs

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/

// Approach: Dijkstra's
// Time: O(mnlog(mn))
// Space: O(mn)
func minimumTime2577(grid [][]int) int {
	if min(grid[0][1], grid[1][0]) > 1 {
		return -1
	}
	type Cell struct {
		Row  int
		Col  int
		Time int
	}
	var abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	var directions = [5]int{-1, 0, 1, 0, -1}
	var minHeap = pkg.NewHeap(
		[]*Cell{{0, 0, 0}},
		func(x, y *Cell) bool {
			return x.Time < y.Time
		},
	)
	var m, n = len(grid), len(grid[0])
	var seen = make([][]bool, m)
	for idx := range seen {
		seen[idx] = make([]bool, n)
	}
	seen[0][0] = true
	var row, col, wait, time int
	for minHeap.Len() > 0 {
		var curr = heap.Pop(minHeap).(*Cell)
		if curr.Row == m-1 && curr.Col == n-1 {
			return curr.Time
		}
		for dir := 1; dir < 5; dir++ {
			row = curr.Row + directions[dir-1]
			col = curr.Col + directions[dir]
			if min(row, col) < 0 || row == m || col == n || seen[row][col] {
				continue
			}
			wait = 0
			if abs(grid[row][col]-curr.Time)&1 == 0 {
				wait = 1
			}
			time = max(curr.Time+1, grid[row][col]+wait)
			heap.Push(minHeap, &Cell{row, col, time})
			seen[row][col] = true
		}
	}
	return -1
}
