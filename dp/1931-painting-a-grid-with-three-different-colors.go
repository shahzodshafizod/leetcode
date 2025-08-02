package dp

// https://leetcode.com/problems/painting-a-grid-with-three-different-colors/

// Approach: Top-Down Dynamic Programming
// Time: O(m*n*2^10)
// Space: O(m*n*2^10)
func colorTheGrid(m int, n int) int {
	const mod int = 1e9 + 7

	memo := make([][][1023]*int, m) // 1023 = (1<<10)-1
	for idx := range memo {
		memo[idx] = make([][1023]*int, n)
	}

	shiftpos := (m - 1) * 2

	var dp func(mask int, row int, col int) int

	dp = func(mask int, row int, col int) int {
		if row == m {
			return dp(mask, 0, col+1)
		}

		if col == n {
			return 1
		}

		if memo[row][col][mask] != nil {
			return *memo[row][col][mask]
		}

		up := 0
		if row > 0 {
			up = mask >> shiftpos
		}

		left := 0
		if col > 0 {
			left = mask & 0b11
		}

		newMask := mask >> 2
		res := 0

		for color := 1; color <= 3; color++ {
			if color == up || color == left {
				continue
			}

			res = (res + dp(newMask|(color<<shiftpos), row+1, col)) % mod
		}

		memo[row][col][mask] = &res

		return res
	}

	return dp(0, 0, 0)
}
