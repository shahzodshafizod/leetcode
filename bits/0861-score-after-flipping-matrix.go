package bits

/*
There are only two rules that you need to know in this problem:
	If the first number in the row is 0, flip the row.
	If the count of 0 in the col is greater than the count of 1, flip the col.
*/

// https://leetcode.com/problems/score-after-flipping-matrix/

// time: O(m x n)
// space: O(1)
func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 1. all bits in the first column should be set
	score := m * (1 << (n - 1))

	var count int
	for col := 1; col < n; col++ {
		count = 0

		for row := range m {
			/*
				For counting 0's, the first and the current should be different:
					2.1. if the first element of this row is 1,
						then we check if the current is 0
					2.2. if the first element of this row is 0,
						then we know the current should be flipped,
						so, we check if the current is 1
				For counting 1's, just check if the first and the current are equal
			*/
			if grid[row][col] != grid[row][0] {
				count++
			}
		}

		score += max(count, m-count) * (1 << (n - 1 - col))
	}

	return score
}

// // time: O(m x n)
// // space: O(1)
// func matrixScore(grid [][]int) int {
// 	var m, n = len(grid), len(grid[0])
// 	for row := 0; row < m; row++ { // O(m x n)
// 		if grid[row][0] == 0 {
// 			for col := 0; col < n; col++ {
// 				grid[row][col] ^= 1
// 			}
// 		}
// 	}
// 	var score = 0
// 	var zeroes int
// 	for col := 0; col < n; col++ { // O(n x m)
// 		zeroes = 0
// 		for row := 0; row < m; row++ {
// 			if grid[row][col] == 0 {
// 				zeroes++
// 			}
// 		}
// 		score += max(m-zeroes, zeroes) * (1 << (n - 1 - col))
// 	}
// 	return score
// }
