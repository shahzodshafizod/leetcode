package matrices

// https://leetcode.com/problems/snakes-and-ladders/

// Approach: Breadth-First Search
// Time: O(nn)
// Space: O(nn)
func snakesAndLadders(board [][]int) int {
	var n = len(board)
	var getPos = func(square int) (int, int) {
		var row = (square - 1) / n
		var col = (square - 1) % n
		if row&1 == 1 {
			col = n - 1 - col
		}
		return n - 1 - row, col
	}
	var target = n * n
	var queue = make([][2]int, target)
	queue[0] = [2]int{1, 0} // [2]int{square, moves}
	var qid, qlen = 0, 1
	var visited = make([]bool, target)
	var square, moves, row, col, next int
	for qid < qlen {
		square, moves = queue[qid][0], queue[qid][1]+1
		qid++
		for i := min(square+6, target); i > square; i-- {
			next = i
			row, col = getPos(next)
			if board[row][col] != -1 {
				next = board[row][col]
			}
			if next == target {
				return moves
			}
			if !visited[next-1] {
				visited[next-1] = true
				queue[qlen][0] = next
				queue[qlen][1] = moves
				qlen++
			}
		}
	}
	return -1
}
