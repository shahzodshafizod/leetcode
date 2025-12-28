package matrices

// https://leetcode.com/problems/image-smoother/

// Approach: Check 3x3 grid for each cell and calculate average
// For each cell, iterate through all valid neighbors in 3x3 grid
// and calculate the average (floor division)
// Time: O(m*n) - visit each cell once, check constant 9 neighbors
// Space: O(m*n) - result matrix (required by problem)
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])

	result := make([][]int, m)
	for i := range m {
		result[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			totalSum := 0
			count := 0

			// Check 3x3 grid centered at (i, j)
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					ni, nj := i+di, j+dj
					if ni >= 0 && ni < m && nj >= 0 && nj < n {
						totalSum += img[ni][nj]
						count++
					}
				}
			}

			result[i][j] = totalSum / count
		}
	}

	return result
}
