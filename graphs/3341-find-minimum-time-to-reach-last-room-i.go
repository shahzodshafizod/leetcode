package graphs

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/

// Approach: Shortest Path + Dijkstra
// Time: O(nm log nm)
// Space: O(nm)
func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	seen := make([][]bool, n)

	for row := 0; row < n; row++ {
		seen[row] = make([]bool, m)
	}

	minHeap := pkg.NewHeap(
		[][3]int{{0, 0, 0}},
		func(x, y [3]int) bool { return x[2] < y[2] },
	)
	seen[0][0] = true
	directions := [5]int{1, 0, -1, 0, 1}

	var row, col, time, newR, newC, newT int

	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).([3]int)

		row, col, time = top[0], top[1], top[2]
		for dir := 1; dir < 5; dir++ {
			newR = row + directions[dir-1]
			newC = col + directions[dir]

			if min(newR, newC) >= 0 && newR < n && newC < m && !seen[newR][newC] {
				newT = max(time, moveTime[newR][newC]) + 1
				if newR == n-1 && newC == m-1 {
					return newT
				}

				heap.Push(minHeap, [3]int{newR, newC, newT})

				seen[newR][newC] = true
			}
		}
	}

	return -1
}
