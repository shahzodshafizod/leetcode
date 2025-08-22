package monotonic

// https://leetcode.com/problems/count-submatrices-with-all-ones/

// Approach #4: Monotonic Stack
// Time: O(mxn)
// Space: O(n)
func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])

	var count int

	height := make([]int, n)
	for row := range m {
		subRes := make([]int, n)
		stack := make([]int, n)
		size := 0

		for col := range n {
			if mat[row][col] == 1 {
				height[col]++
			} else {
				height[col] = 0
			}

			for size > 0 && height[stack[size-1]] >= height[col] {
				size--
			}

			if size > 0 {
				prev := stack[size-1]
				subRes[col] = subRes[prev] + height[col]*(col-prev)
			} else {
				subRes[col] = height[col] * (col + 1)
			}

			stack[size] = col
			size++
			count += subRes[col]
		}
	}

	return count
}

// // Approach #3: Enumeration
// // Time: O(mmxn)
// // Space: O(1)
// func numSubmat(mat [][]int) int {
// 	m, n := len(mat), len(mat[0])

// 	var count int

// 	for row := range m {
// 		for col := range n {
// 			if mat[row][col] == 1 {
// 				if col > 0 {
// 					mat[row][col] += mat[row][col-1]
// 				}
// 			}

// 			cur := mat[row][col]
// 			for k := row; k >= 0; k-- {
// 				cur = min(cur, mat[k][col])
// 				if cur == 0 {
// 					break
// 				}

// 				count += cur
// 			}
// 		}
// 	}

// 	return count
// }

// // Approach #2: Enumeration
// // Time: O(mxnn)
// // Space: O(n)
// func numSubmat(mat [][]int) int {
// 	m, n := len(mat), len(mat[0])

// 	var count, minHeight int

// 	height := make([]int, n)
// 	for row := range m {
// 		// calculate height for the current row
// 		for col := range n {
// 			if mat[row][col] == 0 {
// 				height[col] = 0
// 			} else {
// 				height[col]++
// 			}
// 		}
// 		// for each ceil, count rectangles ending at (row, col)
// 		for col := range n {
// 			if height[col] > 0 {
// 				minHeight = height[col]
// 				for k := col; k >= 0; k-- {
// 					minHeight = min(minHeight, height[k])
// 					if minHeight == 0 {
// 						break
// 					}

// 					count += minHeight
// 				}
// 			}
// 		}
// 	}

// 	return count
// }

// // Approach #1: Brute-Force
// // Time: O(nnxmm)
// // Space: O(1)
// func numSubmat(mat [][]int) int {
// 	var count, bound int

// 	m, n := len(mat), len(mat[0])
// 	for row := range m {
// 		for col := range n {
// 			bound = n
// 			for i := row; i < m; i++ {
// 				for j := col; j < bound; j++ {
// 					if mat[i][j] == 1 {
// 						count++
// 					} else {
// 						bound = j
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return count
// }
