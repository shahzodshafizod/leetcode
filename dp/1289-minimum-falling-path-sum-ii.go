package dp

// https://leetcode.com/problems/minimum-falling-path-sum-ii/

// Tabulation
// time: O(NxN)
// space: O(1)
func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	var min1col, min2col int
	for row := 1; row < n; row++ {
		// 1. find two minimums (min1<min2) from prev row
		min1col, min2col = 0, 1
		if grid[row-1][min1col] > grid[row-1][min2col] {
			min1col, min2col = min2col, min1col
		}
		for col := 2; col < n; col++ {
			if grid[row-1][col] < grid[row-1][min1col] {
				min2col = min1col
				min1col = col
			} else if grid[row-1][col] < grid[row-1][min2col] {
				min2col = col
			}
		}
		// 2. replace min1<-min2, others<-min1 in the current row
		for col := 0; col < n; col++ {
			if col == min1col {
				grid[row][col] += grid[row-1][min2col]
			} else {
				grid[row][col] += grid[row-1][min1col]
			}
		}
	}
	// 3. find the min from the last row
	minsum := grid[n-1][0]
	for col := 1; col < n; col++ {
		minsum = min(minsum, grid[n-1][col])
	}
	return minsum
}

// // Memoization
// // time: O(NxN)
// // space: O(NxN)
// func minFallingPathSum(grid [][]int) int {
// 	var m, n = len(grid), len(grid[0])
// 	var cache = make([][]*int, m)
// 	for idx := range cache {
// 		cache[idx] = make([]*int, n)
// 	}
// 	var dp func(row int, col int) int
// 	dp = func(row int, col int) int {
// 		if cache[row][col] != nil {
// 			return *cache[row][col]
// 		}
// 		var minsum = grid[row][col]
// 		var inited = false
// 		var sum int
// 		if row+1 < m {
// 			for c := 0; c < n; c++ {
// 				if c == col {
// 					continue
// 				}
// 				sum = grid[row][col] + dp(row+1, c)
// 				if !inited || sum < minsum {
// 					minsum = sum
// 					inited = true
// 				}
// 			}
// 		}
// 		cache[row][col] = &minsum
// 		return minsum
// 	}
// 	var sum = math.MaxInt
// 	for col := 0; col < n; col++ {
// 		sum = min(sum, dp(0, col))
// 	}
// 	return sum
// }
