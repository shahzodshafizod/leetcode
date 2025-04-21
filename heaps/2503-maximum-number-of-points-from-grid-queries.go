package heaps

import (
	"container/heap"
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/

// Approach: Sorting Queries + Min-Heap Expansion
// Time: O(klogk+n⋅mlog(n⋅m))
// Space: O(n⋅m+k)
func maxPoints(grid [][]int, queries []int) []int {
	var qn = len(queries)
	var qindices = make([][2]int, qn)
	for idx := 0; idx < qn; idx++ {
		qindices[idx][0] = queries[idx]
		qindices[idx][1] = idx
	}
	sort.Slice(qindices, func(i, j int) bool {
		return qindices[i][0] < qindices[j][0]
	})
	var m, n = len(grid), len(grid[0])
	var seen = make([][]bool, m)
	for row := 0; row < m; row++ {
		seen[row] = make([]bool, n)
	}
	var queue = pkg.NewHeap(
		[][2]int{{0, 0}},
		func(x, y [2]int) bool {
			return grid[x[0]][x[1]] < grid[y[0]][y[1]]
		},
	)
	seen[0][0] = true
	var directions = [5]int{1, 0, -1, 0, 1}
	var answer = make([]int, qn)
	var row, col, nr, nc int
	var top [2]int
	var points = 0
	for idx := 0; idx < qn; idx++ {
		for queue.Len() > 0 && grid[queue.Peak()[0]][queue.Peak()[1]] < qindices[idx][0] {
			top = heap.Pop(queue).([2]int)
			row, col = top[0], top[1]
			for dir := 1; dir < 5; dir++ {
				nr = row + directions[dir-1]
				nc = col + directions[dir]
				if min(nr, nc) >= 0 && nr < m && nc < n && !seen[nr][nc] {
					heap.Push(queue, [2]int{nr, nc})
					seen[nr][nc] = true
				}
			}
			points++
		}
		answer[qindices[idx][1]] = points
	}
	return answer
}
