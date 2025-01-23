package matrices

// https://leetcode.com/problems/count-servers-that-communicate/

// Approach: Counting
// Time: O(mn)
// Space: O(m+n)
func countServers(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	var rcount, ccount = make([]int, m), make([]int, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			rcount[row] += grid[row][col]
			ccount[col] += grid[row][col]
		}
	}
	var count = 0
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == 1 && rcount[row]+ccount[col] > 2 {
				count++
			}
		}
	}
	return count
}
