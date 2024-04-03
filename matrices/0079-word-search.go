package matrices

// https://leetcode.com/problems/word-search/

// time: O(m*n*4^w)
// space: O(2*w) = O(w); O(w)-recursion stack, O(w)-counts
func exist(board [][]byte, word string) bool {
	var m, n, w = len(board), len(board[0]), len(word)
	var letters = []byte(word)
	var counts = make([]int, w)
	for row := 0; row < m; row++ { // time: O(m*n*w)
		for col := 0; col < n; col++ {
			for idx := range letters {
				if board[row][col] == letters[idx] {
					counts[idx]++
				}
			}
		}
	}
	// Follow up: answer#1
	for _, count := range counts {
		if count == 0 {
			return false
		}
	}
	// Follow up: answer#2
	if counts[0] > counts[w-1] {
		for left, right := 0, w-1; left < right; { // O(w/2) = O(w)
			letters[left], letters[right] = letters[right], letters[left]
			left++
			right--
		}
	}
	var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(idx int, row int, col int) bool
	dfs = func(idx int, row int, col int) bool {
		if idx == w {
			return true
		}
		if row < 0 || row == m || col < 0 || col == n {
			return false
		}
		letter := board[row][col]
		if letter == '.' || letter != letters[idx] {
			return false
		}
		board[row][col] = '.'
		for _, direction := range directions {
			r, c := row+direction[0], col+direction[1]
			if dfs(idx+1, r, c) {
				board[row][col] = letter
				return true
			}
		}
		// backtracking
		board[row][col] = letter
		return false
	}
	for row := 0; row < m; row++ { // O(m*n*4^w)
		for col := 0; col < n; col++ {
			if dfs(0, row, col) { // O(4^w)
				return true
			}
		}
	}
	return false
}

/*
Follow up: Could you use search pruning to make your solution faster with a larger board?

#1.	You can also use the frequency counting of the letters of the word in the board
	to check if some letter of the word doesn't exist in the board,
	returning false immediately without ever backtraking.
	This was the pruning that get me out of the TLE.

#2. And also a smart pruning technique is to check the frequency
	of first and last char of the word in the board,
	and if the freq of first char is greater than last one,
	simple reverse the string and find the result.
	It will reduce the number of starting cells for dfs.
*/
