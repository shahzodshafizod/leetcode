package matrices

// https://leetcode.com/problems/map-of-highest-peak/

// Approach: Breadth-First Search
// Time: O(mn)
// Space: O(mn)
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	queue := make([][2]int, 0)

	height := make([][]int, m)
	for row := range m {
		height[row] = make([]int, n)

		for col := range n {
			if isWater[row][col] == 1 {
				queue = append(queue, [2]int{row, col})
			} else {
				height[row][col] = -1
			}
		}
	}

	directions := [5]int{-1, 0, 1, 0, -1}

	var row, col, r, c int

	for len(queue) > 0 {
		next := make([][2]int, 0)

		for idx := range queue {
			row, col = queue[idx][0], queue[idx][1]
			for dir := 1; dir < 5; dir++ {
				r, c = row+directions[dir-1], col+directions[dir]
				if min(r, c) < 0 || r == m || c == n || height[r][c] != -1 {
					continue
				}

				height[r][c] = height[row][col] + 1

				next = append(next, [2]int{r, c})
			}
		}

		queue = next
	}

	return height
}
