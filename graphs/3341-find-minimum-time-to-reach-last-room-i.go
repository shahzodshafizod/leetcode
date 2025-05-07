package graphs

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/

// Approach: Shortest Path + Dijkstra
// Time: O(nmlog(nm))
// Space: O(nm)
func minTimeToReach(moveTime [][]int) int {
	var m, n = len(moveTime), len(moveTime[0])
	var seen = make([][]bool, m)
	for row := 0; row < m; row++ {
		seen[row] = make([]bool, n)
	}
	var minHeap = pkg.NewHeap(
		[][3]int{{0, 0, 0}},
		func(x, y [3]int) bool { return x[2] < y[2] },
	)
	seen[0][0] = true
	var directions = [5]int{1, 0, -1, 0, 1}
	var row, col, time, newR, newC int
	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).([3]int)
		row, col, time = top[0], top[1], top[2]
		if row == m-1 && col == n-1 {
			break
		}
		for dir := 1; dir < 5; dir++ {
			newR = row + directions[dir-1]
			newC = col + directions[dir]
			if min(newR, newC) >= 0 && newR < m && newC < n && !seen[newR][newC] {
				heap.Push(minHeap, [3]int{newR, newC, max(time, moveTime[newR][newC]) + 1})
				seen[newR][newC] = true
			}
		}
	}
	return time
}
