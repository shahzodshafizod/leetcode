package unionfinds

// https://leetcode.com/problems/last-day-where-you-can-still-cross/

// Approach #2: Union Find
// Time: O(row * col)
// Space: O(row * col)
func latestDayToCross(row int, col int, cells [][]int) int {
	n := row*col + 2

	root := make([]int, n)
	for i := range n {
		root[i] = i
	}

	var find func(x int) int

	find = func(x int) int {
		if root[x] != x {
			root[x] = find(root[x])
		}

		return root[x]
	}
	union := func(x int, y int) {
		rx, ry := find(x), find(y)
		if rx < ry {
			root[ry] = rx
		} else if ry < rx {
			root[rx] = ry
		}
	}

	grid := make([][]int, row)
	for i := range row {
		grid[i] = make([]int, col)
	}

	directions := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

	var r, c, nr, nc, curr int

	left, right := 0, n-1

	var day int

	for i := range cells {
		r, c = cells[i][0]-1, cells[i][1]-1
		grid[r][c] = 1

		curr = r*col + c + 1
		for d := range directions {
			nr, nc = r+directions[d][0], c+directions[d][1]
			if min(nr, nc) >= 0 && nr < row && nc < col && grid[nr][nc] == 1 {
				union(curr, nr*col+nc+1)
			}
		}

		switch c {
		case 0:
			union(curr, left)
		case col - 1:
			union(curr, right)
		default:
		}

		if find(left) == find(right) {
			break
		}

		day++
	}

	return day
}

// // Approach #1: Binary Search, DFS
// // Time: O(row * col * log(len(cells)))
// // Space: O(row * col)
// func latestDayToCross(row int, col int, cells [][]int) int {
// 	directions := [5]int{0, 1, 0, -1, 0}

// 	var dfs func(seen [][]bool, r int, c int) bool

// 	dfs = func(seen [][]bool, r int, c int) bool {
// 		if min(r, c) < 0 || r == row || c == col || seen[r][c] {
// 			return false
// 		}

// 		if r == row-1 {
// 			return true
// 		}

// 		seen[r][c] = true

// 		var nr, nc int
// 		for d := range 4 {
// 			nr = r + directions[d]

// 			nc = c + directions[d+1]
// 			if dfs(seen, nr, nc) {
// 				return true
// 			}
// 		}

// 		return false
// 	}

// 	var (
// 		mid, r, c int
// 		possible  bool
// 	)

// 	left, right := 0, len(cells)
// 	for left <= right {
// 		mid = left + (right-left)/2

// 		seen := make([][]bool, row)
// 		for i := range seen {
// 			seen[i] = make([]bool, col)
// 		}

// 		for i := range mid {
// 			r, c = cells[i][0], cells[i][1]
// 			seen[r-1][c-1] = true
// 		}

// 		possible = false

// 		for c := range col {
// 			if dfs(seen, 0, c) {
// 				possible = true

// 				break
// 			}
// 		}

// 		if possible {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	return right
// }
