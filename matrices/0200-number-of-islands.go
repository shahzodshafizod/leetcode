package matrices

/*
2-D Arrays
Approaches:
	- Sequential
	- BFS
	- DFS

Problem:
Given a 2D array containing only 1's (land) and 0's (water), count the number of islands.
An island is land connected hotizontally or vertically.

Step 1: Verify the constraints
	- Are the edges of the 2D array water?
		: Yes, assume anything outside of the 2D array is water.
Step 2: Write out some test cases
	- [['1', '1', '1', '1', '0'],		-> 2
	   ['1', '1', '0', '1', '0'],
	   ['1', '1', '0', '0', '1'],
	   ['0', '0', '0', '1', '1']]

	- [['0', '1', '0', '1', '0'],		-> 7
	   ['1', '0', '1', '0', '1'],
	   ['0', '1', '1', '1', '0'],
	   ['1', '0', '1', '0', '1']]

	- []								-> 0

	- [[], []]							-> 0

	- [['0', '0', '0', '0', '0'],		-> 0
	   ['0', '0', '0', '0', '0'],
	   ['0', '0', '0', '0', '0'],
	   ['0', '0', '0', '0', '0']]

	- [['1', '1', '1', '1', '1'],		-> 1
	   ['1', '1', '1', '1', '1'],
	   ['1', '1', '1', '1', '1'],
	   ['1', '1', '1', '1', '1']]
*/

// https://leetcode.com/problems/number-of-islands/

// DFS
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var drowning func(row int, col int)

	drowning = func(row int, col int) {
		if grid[row][col] == '0' {
			return
		}

		grid[row][col] = '0'

		if row+1 < m { // south
			drowning(row+1, col)
		}

		if row-1 >= 0 { // north
			drowning(row-1, col)
		}

		if col+1 < n { // east
			drowning(row, col+1)
		}

		if col-1 >= 0 { // west
			drowning(row, col-1)
		}
	}
	num := 0

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == '1' {
				num++

				drowning(row, col)
			}
		}
	}

	return num
}
