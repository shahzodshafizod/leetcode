package matrices

/*
Problem:
Given a 2D array containing 0's (empty cell), 1's (fresh oranges) and 2's (rotten oranges).
Every minute, all fresh oranges immediately adjacent (4 directions) to rotten oranges will rot.
How many minutes must pass until all oranges are rotten?

Step 1: Verify the constraints
	- What do we return if it's not possible?
		: If all the oranges can't rot, return -1.
	- What do we return if there are no oranges?
		: Return 0 if there are no oranges.
Step 2: Write out some test cases
	- [[2, 1, 1, 0, 0],		-> 7
	   [1, 1, 1, 0, 0],
	   [0, 1, 1, 1, 1],
	   [0, 1, 0, 0, 1]]

	- [[1, 1, 0, 0, 0],		-> -1
	   [2, 1, 0, 0, 0],
	   [0, 0, 0, 1, 2],
	   [0, 1, 0, 0, 1]]

Approach:
	- sequential order:
		- count fresh oranges
		- put rotten oranges into queue
	- BFS:
		- use queue size to track minutes
		- process rotting oranges:
			- rot adjacent fresh oranges
			- push into queue
			- decrement fresh oranges count
*/

// https://leetcode.com/problems/rotting-oranges/

func orangesRotting(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}

	n := len(grid[0])
	if n == 0 {
		return 0
	}

	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down and left
	rottens := make([][2]int, 0)
	freshes := 0

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			switch grid[row][col] {
			case 1:
				freshes++
			case 2:
				rottens = append(rottens, [2]int{row, col})
			}
		}
	}

	minutes := 0

	for length := len(rottens); length > 0; {
		for i := 0; i < length; i++ {
			for _, direction := range directions {
				row, col := rottens[i][0]+direction[0], rottens[i][1]+direction[1]
				if row >= 0 && col >= 0 && row < m && col < n && grid[row][col] == 1 {
					grid[row][col] = 2
					freshes--

					rottens = append(rottens, [2]int{row, col})
				}
			}
		}

		rottens = rottens[length:]

		length = len(rottens)
		if length > 0 {
			minutes++
		}
	}

	if freshes > 0 {
		minutes = -1
	}

	return minutes
}
