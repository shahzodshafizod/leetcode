package dp

// https://leetcode.com/problems/unique-paths-ii/

// time: O(M*N)
// space: O(N)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	var memo = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if obstacleGrid[m-1][i] == 1 {
			break
		}
		memo[i] = 1
	}
	for row := m - 2; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			if obstacleGrid[row][col] == 1 {
				memo[col] = 0
			} else if col < n-1 && obstacleGrid[row][col+1] == 0 {
				memo[col] += memo[col+1]
			}
		}
	}
	return memo[0]
}

// // time: O(M*N)
// // space: O(M*N) (recursive stack)
// func uniquePathsWithObstacles(obstacleGrid [][]int) int {
// 	m := len(obstacleGrid)
// 	n := len(obstacleGrid[0])
// 	for row := 0; row < m; row++ {
// 		for col := 0; col < n; col++ {
// 			if obstacleGrid[row][col] == 1 {
// 				obstacleGrid[row][col] = -1
// 			}
// 		}
// 	}
// 	var dp func(int, int) int
// 	dp = func(row int, col int) int {
// 		if row == m || col == n {
// 			return 0
// 		}
// 		if obstacleGrid[row][col] == -1 {
// 			return 0
// 		}
// 		if row == m-1 && col == n-1 {
// 			return 1
// 		}
// 		if obstacleGrid[row][col] > 0 {
// 			return obstacleGrid[row][col]
// 		}
// 		obstacleGrid[row][col] = dp(row+1, col) + dp(row, col+1)
// 		return obstacleGrid[row][col]
// 	}
// 	return dp(0, 0)
// }
