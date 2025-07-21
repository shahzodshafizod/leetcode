package tries

// https://leetcode.com/problems/word-search-ii/

func findWords(board [][]byte, words []string) []string {
	type node struct { // trie node
		word     string
		children [26]*node
	}
	insert := func(curr *node, word *string) {
		for _, c := range *word {
			c -= 'a'
			if curr.children[c] == nil {
				curr.children[c] = &node{}
			}
			curr = curr.children[c]
		}
		curr.word = *word
	}
	m, n := len(board), len(board[0])
	var dfs func(row int, col int, curr *node, result *[]string)
	dfs = func(row int, col int, curr *node, result *[]string) {
		letter := board[row][col]
		if letter == '.' || curr.children[letter-'a'] == nil {
			return
		}
		next := curr.children[letter-'a']
		if next.word != "" {
			*result = append(*result, next.word)
			next.word = ""
		}
		board[row][col] = '.'
		if row > 0 {
			dfs(row-1, col, next, result) // north
		}
		if row+1 < m {
			dfs(row+1, col, next, result) // south
		}
		if col > 0 {
			dfs(row, col-1, next, result) // west
		}
		if col+1 < n {
			dfs(row, col+1, next, result) // east
		}
		board[row][col] = letter
	}

	root := &node{}
	for _, word := range words {
		insert(root, &word)
	}
	result := make([]string, 0)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			dfs(row, col, root, &result)
		}
	}
	return result
}
