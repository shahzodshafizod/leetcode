package graphs

// https://leetcode.com/problems/word-ladder/

// Approach: Breadth-First Search
// N = len(wordList)
// M = len(wordList[i])
// Time: O(N * M^2)
// Space: O(N*M)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	adjList := make(map[string][]string)
	exists := false
	var pattern []byte
	var c byte
	wordList = append(wordList, beginWord)
	for _, word := range wordList {
		exists = exists || word == endWord
		pattern = []byte(word)
		for idx := range pattern {
			c = pattern[idx]
			pattern[idx] = '*'
			adjList[string(pattern)] = append(adjList[string(pattern)], word)
			pattern[idx] = c
		}
	}
	if !exists {
		return 0
	}
	n := len(wordList)
	queue := make([]string, n)
	queue[0] = beginWord
	head, tail := 0, 0
	visited := make(map[string]bool)
	visited[beginWord] = true
	steps := 0
	for head <= tail {
		steps++
		end := tail
		for head <= end {
			word := queue[head]
			head++
			if word == endWord {
				return steps
			}
			pattern = []byte(word)
			for idx := range pattern {
				c = pattern[idx]
				pattern[idx] = '*'
				for _, next := range adjList[string(pattern)] {
					if !visited[next] {
						visited[next] = true
						tail++
						queue[tail] = next
					}
				}
				pattern[idx] = c
			}
		}
	}
	return 0
}
