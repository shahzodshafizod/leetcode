package matrices

import (
	"container/list"
	"math"
)

// https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/

// Approach: Breadth-First Search
// Time: O(mxn)
// Space: O(mxn)
func minCost(grid [][]int) int {
	directions := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	minCost := make([][]int, m)

	for row := range minCost {
		minCost[row] = make([]int, n)
		for col := range minCost[row] {
			minCost[row][col] = math.MaxInt
		}
	}

	queue := list.New()
	queue.PushBack([3]int{0, 0, 0}) // row, col, cost

	var row, col, cost, r, c int

	for queue.Len() > 0 {
		item := queue.Remove(queue.Front()).([3]int)

		row, col, cost = item[0], item[1], item[2]
		if min(row, col) < 0 || row == m || col == n || cost >= minCost[row][col] {
			continue
		}

		minCost[row][col] = cost

		if row == m-1 && col == n-1 {
			break
		}

		for dir := range directions {
			r = row + directions[dir][0]
			c = col + directions[dir][1]

			if dir+1 == grid[row][col] {
				queue.PushFront([3]int{r, c, cost})
			} else {
				queue.PushBack([3]int{r, c, cost + 1})
			}
		}
	}

	return minCost[m-1][n-1]
}
