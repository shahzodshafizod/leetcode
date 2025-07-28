package graphs

// https://leetcode.com/problems/word-ladder-ii/

// Approach: Breadth-First Search + Backtracking
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	m := len(beginWord)
	graph := make(map[string][]string)

	found := false
	for _, word := range wordList {
		found = found || word == endWord

		pattern := []byte(word)
		for idx := 0; idx < m; idx++ {
			tmp := pattern[idx]
			pattern[idx] = '*'
			graph[string(pattern)] = append(graph[string(pattern)], word)
			pattern[idx] = tmp
		}
	}

	if !found {
		return [][]string{}
	}

	prev := make(map[string][]string)
	queue := []string{beginWord}
	visited := map[string]int{beginWord: 1}
	found = false
	layer := 0

	for size := len(queue); size > 0 && !found; size = len(queue) {
		layer++

		for _, word := range queue {
			pattern := []byte(word)
			for idx := 0; idx < m; idx++ {
				tmp := pattern[idx]

				pattern[idx] = '*'
				for _, next := range graph[string(pattern)] {
					if _, exists := visited[next]; !exists || visited[next] == layer+1 {
						found = found || next == endWord
						prev[next] = append(prev[next], word)

						if !exists {
							visited[next] = layer + 1

							queue = append(queue, next)
						}
					}
				}

				pattern[idx] = tmp
			}
		}

		queue = queue[size:]
	}
	// backtracking
	var dfs func(word string, path []string) [][]string

	dfs = func(word string, path []string) [][]string {
		if word == beginWord {
			return [][]string{path}
		}

		sequences := make([][]string, 0)

		for _, prev := range prev[word] {
			newPath := append([]string{prev}, path...)

			seq := dfs(prev, newPath)
			if len(seq) > 0 {
				sequences = append(sequences, seq...)
			}
		}

		return sequences
	}

	return dfs(endWord, []string{endWord})
}
