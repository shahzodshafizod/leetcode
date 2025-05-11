package graphs

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-minimum-time-to-reach-last-room-ii/

// Approach: Shortest Path + Dijkstra
// Time: O(nm log nm)
// Space: O(nm)
func minTimeToReach3342(moveTime [][]int) int {
	var n, m = len(moveTime), len(moveTime[0])
	var seen = make([][]bool, n)
	for row := 0; row < n; row++ {
		seen[row] = make([]bool, m)
	}
	var minheap = pkg.NewHeap(
		[][3]int{{0, 0, 0}}, // {row, col, time}
		func(x, y [3]int) bool { return x[2] < y[2] },
	)
	var directions = [5]int{1, 0, -1, 0, 1}
	var row, col, time, newr, newc, newt, step int
	for minheap.Len() > 0 {
		top := heap.Pop(minheap).([3]int)
		row, col, time = top[0], top[1], top[2]
		step = (row+col)%2 + 1
		for dir := 1; dir < 5; dir++ {
			newr = row + directions[dir-1]
			newc = col + directions[dir]
			if min(newr, newc) >= 0 && newr < n && newc < m && !seen[newr][newc] {
				newt = max(time, moveTime[newr][newc]) + step
				if newr == n-1 && newc == m-1 {
					return newt
				}
				heap.Push(minheap, [3]int{newr, newc, newt})
				seen[newr][newc] = true
			}
		}
	}
	return -1
}
