package dp

// https://leetcode.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	var dp func(row int, col int) int
	dp = func(row int, col int) int {
		if row == m || col == n {
			return -1
		}
		if grid[row][col]>>8 > 0 {
			return grid[row][col] >> 8
		}
		var sum = grid[row][col]
		down := dp(row+1, col)
		right := dp(row, col+1)
		if down != -1 && right != -1 {
			sum += min(down, right)
		} else if down != -1 {
			sum += down
		} else if right != -1 {
			sum += right
		}
		grid[row][col] += sum << 8
		return sum
	}
	res := dp(0, 0)
	return res
}
