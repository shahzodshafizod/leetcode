package dp

// https://leetcode.com/problems/find-the-maximum-number-of-fruits-collected/

// Approach: Top-Down Dynamic Programming
// Time: O(nn)
// Space: O(nn)
func maxCollectedFruits(fruits [][]int) int {
	n := len(fruits)

	memo := make([][]*int, n)

	first := 0
	for i := range n {
		first += fruits[i][i]
		memo[i] = make([]*int, n)
	}

	var dp func(row int, col int, dirs [3][2]int, f int) int

	dp = func(row int, col int, dirs [3][2]int, f int) int {
		if min(row, col) < 0 || max(row, col) == n || f*row <= f*col {
			return 0
		}

		if memo[row][col] != nil {
			return *memo[row][col]
		}

		res := 0
		for _, d := range dirs {
			res = max(res, dp(row+d[0], col+d[1], dirs, f))
		}

		res += fruits[row][col]
		memo[row][col] = &res

		return res
	}

	second := dp(n-1, 0, [3][2]int{{1, 1}, {0, 1}, {-1, 1}}, 1)
	third := dp(0, n-1, [3][2]int{{1, -1}, {1, 0}, {1, 1}}, -1)

	return first + second + third
}
