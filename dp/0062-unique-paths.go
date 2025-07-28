package dp

// https://leetcode.com/problems/unique-paths/

// time: O(M*N)
// space: O(M)
func uniquePaths(m int, n int) int {
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = 1
	}

	for row := m - 1; row > 0; row-- {
		for col := n - 2; col >= 0; col-- {
			count[col] += count[col+1]
		}
	}
	/*
		[84][56][35][20][10][4][1]
		[28][21][15][10][ 6][3][1]
		[ 7][ 6][ 5][ 4][ 3][2][1]
		[ 1][ 1][ 1][ 1][ 1][1][1]
	*/
	return count[0]
}

// // time: O(M*N)
// // space: O(M*N)
// func uniquePaths(m int, n int) int {
// 	var dp func(row int, col int, memo map[int]int) int
// 	dp = func(row int, col int, memo map[int]int) int {
// 		if row == m || col == n {
// 			return 0
// 		}
// 		if row == m-1 && col == n-1 {
// 			return 1
// 		}
// 		key := row*n + col
// 		if memo[key] > 0 {
// 			return memo[key]
// 		}
// 		memo[key] = dp(row+1, col, memo) + dp(row, col+1, memo)
// 		return memo[key]
// 	}
// 	return dp(0, 0, make(map[int]int))
// }

// // 3. combinations
// // space: O((M+N)*max(M,N))
// func uniquePaths(m int, n int) int {
// 	var combinations func(int, int, [][]int) int
// 	combinations = func(n, k int, memo [][]int) int {
// 		if k == 0 || k == n {
// 			return 1
// 		}
// 		if memo[n-1][k-1] == 0 {
// 			memo[n-1][k-1] = combinations(n-1, k, memo) + combinations(n-1, k-1, memo)
// 		}
// 		return memo[n-1][k-1]
// 	}
// 	// m-1 steps down
// 	// n-1 steps right
// 	steps := m + n - 2 // (m-1)+(n-1) = m-1+n-1 = m+n-2
// 	places := m - 1    // n-1 is also valid
// 	var memo = make([][]int, steps)
// 	for i := 0; i < steps; i++ {
// 		memo[i] = make([]int, places)
// 	}
// 	return combinations(steps, places, memo)
// }
