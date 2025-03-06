package matrices

// https://leetcode.com/problems/find-missing-and-repeated-values/

// Approach: Counting
// Time: O(nn)
// Space: O(nn)
func findMissingAndRepeatedValues(grid [][]int) []int {
	var n = len(grid)
	var nn = n * n
	var seen = make([]bool, nn+1)
	var sum = nn * (nn + 1) / 2
	var twice int
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if !seen[grid[row][col]] {
				seen[grid[row][col]] = true
				sum -= grid[row][col]
			} else {
				twice = grid[row][col]
			}
		}
	}
	var missing = sum
	return []int{twice, missing}
}
