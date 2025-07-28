package matrices

/*
There can be at most one lucky number in the matrix.

    |      |
----B------X----> r1
    |      |
    |      |
----Y------A----> r2
	|      |
	V      V
    c2     c1

Suppose we have an integer X in the row r1 and column c1.
The integer X is the minimum in its row and maximum in its column and hence is a lucky number.

Let's say there's another integer Y in the column r2 and column c2
Let's assume that Y is also a lucky number.

Let the integer at (r2, c1) be A
X>A
Y<A
X>A>Y
X>Y

Let the integer at (r1, c2) be B
X<B
Y>B
X<B<Y
X<Y

Hence, the assumption that Y is a lucky number is wrong.
*/

// https://leetcode.com/problems/lucky-numbers-in-a-matrix/

// Approach #2: Greedy
// Time: O(MN)
// Space: O(1)
func luckyNumbers(matrix [][]int) []int {
	const MIN, MAX = 1, 100000

	m, n := len(matrix), len(matrix[0])
	maxOfmins := MIN

	for _, row := range matrix {
		minrow := MAX
		for _, cell := range row {
			minrow = min(minrow, cell)
		}

		maxOfmins = max(maxOfmins, minrow)
	}

	minOfmaxs := MAX

	for col := 0; col < n; col++ {
		maxcol := MIN
		for row := 0; row < m; row++ {
			maxcol = max(maxcol, matrix[row][col])
		}

		minOfmaxs = min(minOfmaxs, maxcol)
	}

	luckies := make([]int, 0)
	if maxOfmins == minOfmaxs {
		luckies = append(luckies, maxOfmins)
	}

	return luckies
}

// // Approach #1: Simulation
// // Time: O(MN)
// // Space: O(M+N)
// func luckyNumbers(matrix [][]int) []int {
// 	var m, n = len(matrix), len(matrix[0])
// 	var mins = make([]int, m)
// 	var maxs = make([]int, n)
// 	for row := 0; row < m; row++ {
// 		mins[row] = matrix[row][0]
// 		for col := 0; col < n; col++ {
// 			mins[row] = min(mins[row], matrix[row][col])
// 			maxs[col] = max(maxs[col], matrix[row][col])
// 		}
// 	}
// 	var luckies = make([]int, 0)
// 	for _, min := range mins {
// 		for _, max := range maxs {
// 			if min == max {
// 				luckies = append(luckies, min)
// 			}
// 		}
// 	}
// 	return luckies
// }
