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

func numIslands(grid [][]byte) int {
	var m = len(grid)
	if m == 0 {
		return 0
	}
	var n = len(grid[0])
	if n == 0 {
		return 0
	}

	var num = 0

	var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down and left
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] != '1' {
				continue
			}

			num++

			// BFS
			var queue = [][2]int{{row, col}}
			grid[row][col] = '0'
			for len(queue) > 0 {
				for _, direction := range directions {
					var r, c = queue[0][0] + direction[0], queue[0][1] + direction[1]
					if r >= 0 && c >= 0 && r < m && c < n && grid[r][c] == '1' {
						queue = append(queue, [2]int{r, c})
						grid[r][c] = '0'
					}
				}
				queue = queue[1:]
			}
		}
	}

	return num
}
