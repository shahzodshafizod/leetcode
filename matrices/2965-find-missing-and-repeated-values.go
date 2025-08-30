package matrices

// https://leetcode.com/problems/find-missing-and-repeated-values/

// Approach: Counting
// Time: O(nn)
// Space: O(nn)
func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	nn := n * n
	seen := make([]bool, nn+1)
	sum := nn * (nn + 1) / 2

	var twice int

	for row := range n {
		for col := range n {
			if !seen[grid[row][col]] {
				seen[grid[row][col]] = true
				sum -= grid[row][col]
			} else {
				twice = grid[row][col]
			}
		}
	}

	missing := sum

	return []int{twice, missing}
}
