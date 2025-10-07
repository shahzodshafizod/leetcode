package matrices

// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/

// Approach: Breadth-First Search
// Time: O(mnk)
// Space: O(mnk)
func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	chance := make([][]int, m)
	for i := range m {
		chance[i] = make([]int, n)
	}

	chance[0][0] = k + 1
	queue := [][3]int{{0, 0, k + 1}}
	directions := [5]int{0, -1, 0, 1, 0}

	var row, col, obs, length, nrow, ncol, nobs int

	for l := len(queue); l > 0; l = len(queue) {
		for i := range l {
			row, col, obs = queue[i][0], queue[i][1], queue[i][2]

			if row == m-1 && col == n-1 {
				return length
			}

			for d := range 4 {
				nrow, ncol = row+directions[d], col+directions[d+1]
				if min(nrow, ncol) < 0 || nrow == m || ncol == n {
					continue
				}

				nobs = obs - grid[nrow][ncol]
				if nobs <= chance[nrow][ncol] {
					continue
				}

				chance[nrow][ncol] = nobs
				queue = append(queue, [3]int{nrow, ncol, nobs})
			}
		}

		length++
		queue = queue[l:]
	}

	return -1
}
