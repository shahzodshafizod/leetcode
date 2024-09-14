package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestLadderLength$
func TestLadderLength(t *testing.T) {
	for _, tc := range []struct {
		beginWord string
		endWord   string
		wordList  []string
		count     int
	}{
		{beginWord: "hit", endWord: "cog", wordList: []string{"hot", "dot", "dog", "lot", "log", "cog"}, count: 5},
		{beginWord: "hit", endWord: "cog", wordList: []string{"hot", "dot", "dog", "lot", "log"}, count: 0},
		{beginWord: "a", endWord: "c", wordList: []string{"a", "b", "c"}, count: 2},
		{beginWord: "leet", endWord: "code", wordList: []string{"lest", "leet", "lose", "code", "lode", "robe", "lost"}, count: 6},
		{beginWord: "hot", endWord: "dog", wordList: []string{"hot", "dog", "dot"}, count: 3},
		{beginWord: "hot", endWord: "dog", wordList: []string{"hot", "dog", "cog", "pot", "dot"}, count: 3},
		{beginWord: "hot", endWord: "dot", wordList: []string{"hot", "dot", "dog"}, count: 2},
		{beginWord: "hot", endWord: "dog", wordList: []string{"hot", "dog"}, count: 0},
		{beginWord: "hbo", endWord: "qbx", wordList: []string{"abo", "hco", "hbw", "ado", "abq", "hcd", "hcj", "hww", "qbq", "qby", "qbz", "qbx", "qbw"}, count: 4},
		{beginWord: "ab", endWord: "ac", wordList: []string{"xx", "ee", "ac"}, count: 2},
		{beginWord: "be", endWord: "ko", wordList: []string{"ce", "mo", "ko", "me", "co"}, count: 4},
	} {
		count := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
		assert.Equal(t, tc.count, count)
	}
}
