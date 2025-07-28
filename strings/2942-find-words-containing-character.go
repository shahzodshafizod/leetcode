package strings

import (
	"strings"
)

// https://leetcode.com/problems/find-words-containing-character/

func findWordsContaining(words []string, x byte) []int {
	var indices []int

	for idx := range words {
		if strings.IndexByte(words[idx], x) != -1 {
			indices = append(indices, idx)
		}
	}

	return indices
}
