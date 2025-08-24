package matrices

// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-ii/

func minimumSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	area := func(lb int, rb int, tb int, db int) int { // lb is left-bound
		left, right := rb+1, lb-1

		top, down := db+1, tb-1
		for row := tb; row <= db; row++ {
			for col := lb; col <= rb; col++ {
				if grid[row][col] == 1 {
					left = min(left, col)
					right = max(right, col)
					top = min(top, row)
					down = max(down, row)
				}
			}
		}

		if left > rb {
			return m*n + 1 // MAX
		}

		return (right - left + 1) * (down - top + 1)
	}

	lb, rb := n, 0

	tb, db := m, 0
	for row := range m {
		for col := range n {
			if grid[row][col] == 1 {
				lb = min(lb, col)
				rb = max(rb, col)
				tb = min(tb, row)
				db = max(db, row)
			}
		}
	}

	res := m * n

	for row := tb; row <= db; row++ {
		for col := lb; col <= rb; col++ {
			// horizontal-up: divide horizontally, up is also divided
			res = min(res,
				area(lb, col, tb, row)+area(col+1, rb, tb, row)+area(lb, rb, row+1, db))

			// horizontal-down: divide horizontally, down is also divided
			res = min(res,
				area(lb, rb, 0, row)+area(lb, col, row+1, db)+area(col+1, rb, row+1, db))

			// vertical-left: divide vertically, left is also divided
			res = min(res,
				area(lb, col, tb, row)+area(lb, col, row+1, db)+area(col+1, rb, tb, db))

			// vertical-right: divide vertically, right is also divided
			res = min(res,
				area(lb, col, tb, db)+area(col+1, rb, tb, row)+area(col+1, rb, row+1, db))
		}
	}

	// two horizontal cuts
	for row1 := range m - 2 {
		for row2 := row1; row2 < m-1; row2++ {
			res = min(res,
				area(lb, rb, tb, row1)+area(lb, rb, row1+1, row2)+area(lb, rb, row2+1, db))
		}
	}

	// two vertical cuts
	for col1 := range n - 2 {
		for col2 := col1; col2 < n-1; col2++ {
			res = min(res,
				area(lb, col1, tb, db)+area(col1+1, col2, tb, db)+area(col2+1, rb, tb, db))
		}
	}

	return res
}
