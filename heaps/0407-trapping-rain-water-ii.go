package heaps

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/trapping-rain-water-ii/

func trapRainWater(heightMap [][]int) int {
	directions := [5]int{-1, 0, 1, 0, -1}
	m, n := len(heightMap), len(heightMap[0])
	queue := pkg.NewHeap(make([][3]int, 0), func(x [3]int, y [3]int) bool { return x[0] < y[0] })
	for row := 0; row < m; row++ {
		heap.Push(queue, [3]int{heightMap[row][0], row, 0})
		heap.Push(queue, [3]int{heightMap[row][n-1], row, n - 1})
		heightMap[row][0] = -1
		heightMap[row][n-1] = -1
	}
	for col := 1; col < n-1; col++ {
		heap.Push(queue, [3]int{heightMap[0][col], 0, col})
		heap.Push(queue, [3]int{heightMap[m-1][col], m - 1, col})
		heightMap[0][col] = -1
		heightMap[m-1][col] = -1
	}
	volume := 0
	var height, row, col, r, c int
	for queue.Len() > 0 {
		top := heap.Pop(queue).([3]int)
		height, row, col = top[0], top[1], top[2]
		for dir := 1; dir < 5; dir++ {
			r, c = row+directions[dir-1], col+directions[dir]
			if min(r, c) < 0 || r == m || c == n || heightMap[r][c] < 0 {
				continue
			}
			if heightMap[r][c] < height {
				volume += height - heightMap[r][c]
				heightMap[r][c] = height
			}
			heap.Push(queue, [3]int{heightMap[r][c], r, c})
			heightMap[r][c] = -1
		}
	}
	return volume
}
