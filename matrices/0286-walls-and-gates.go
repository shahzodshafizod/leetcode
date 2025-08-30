package matrices

// Subscription Required

/*
Problem:
Given a 2D array containing -1's (walls), 0's (gates) and INF's (empty room).
Fill each empty room with the number of steps to the nearest gate.
IF it is impossible to reach a gate, leave INF as the value.
INF is equal to 2147483647.

Step 2: Write out some test cases
	- [[INF, -1, 0, INF],		-> [[3, -1, 0, 1],
	   [INF, INF, INF, -1],			[2, 2, 1, -1],
	   [INF, -1, INF, -1],			[1, -1, 2, -1],
	   [0, -1, INF, INF]]			[0, -1, 3, 4]]

	- [[INF, -1, 0, INF],		-> [[INF, -1, 0, 1],
	   [-1, INF, INF, -1],			[-1, 2, 1, -1],
	   [INF, -1, INF, -1],			[1, -1, 2, -1],
	   [0, -1, INF, INF]]			[0, -1, 3, 4]]

	- []						-> []

	- [[]]						-> [[]]
*/

// https://leetcode.com/problems/walls-and-gates/

func wallsAndGates(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return matrix
	}

	n := len(matrix[0])
	if n == 0 {
		return matrix
	}

	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // up, right, down and left

	for row := range m {
		for col := range n {
			if matrix[row][col] == 0 {
				wallsAndGatesDFS(matrix, directions, m, n, row, col, 0)
			}
		}
	}

	return matrix
}

func wallsAndGatesDFS(
	matrix [][]int,
	directions [4][2]int,
	m int,
	n int,
	row int,
	col int,
	steps int,
) {
	if steps > 0 {
		matrix[row][col] = steps
	}

	steps++

	for _, direction := range directions {
		r, c := row+direction[0], col+direction[1]
		if r >= 0 && c >= 0 && r < m && c < n && matrix[r][c] != -1 && matrix[r][c] != 0 &&
			matrix[r][c] > steps {
			wallsAndGatesDFS(matrix, directions, m, n, r, c, steps)
		}
	}
}
