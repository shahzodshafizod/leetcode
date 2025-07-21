package dp

/*
Problem:
On a given nxn chessboard, a knight piece will start at the r-th row and c-th column.
The knight will attempt to make k moves.
A knight can move in 8 possible ways. Each move will choose one of these 8 at random.
The knight continues moving until it finishes k moves or it moves off the chessboard.
Return the probability that the knight is on the chessboard after it finishes moving.

Step 1: Verify the constraints
	- How many decimals do we round to?
		: Don't round, leave the answer as is.
Step 2: Write out some test cases
	- n: 6, k: 2, row: 2, column: 2			=> 0.53125
	- n: 6, k: 3, row: 2, column: 2			=> 0.359375
	- n: 2, k: 3, row: 1, column: 1			=> 0
	- n: 2, k: 0, row: 1, column: 2			=> 1
*/

// https://leetcode.com/problems/knight-probability-in-chessboard/

func knightProbability(n int, k int, row int, column int) float64 {
	if k == 0 {
		return 1
	}
	prev := make([][]float64, n)
	curr := make([][]float64, n)
	for i := 0; i < n; i++ {
		prev[i] = make([]float64, n)
		curr[i] = make([]float64, n)
	}
	prev[row][column] = 1
	ways := [8][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}}
	var res float64 = 1
	var value float64
	for step := 1; step <= k; step++ {
		res = 0
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				curr[r][c] = 0
				for _, way := range ways {
					prevR, prevC := r+way[0], c+way[1]
					if prevR >= 0 && prevR < n && prevC >= 0 && prevC < n {
						value = prev[prevR][prevC] / 8
						curr[r][c] += value
						res += value
					}
				}
			}
		}
		prev, curr = curr, prev
	}
	return res
}
