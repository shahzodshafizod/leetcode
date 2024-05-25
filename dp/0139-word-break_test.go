package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestWordBreak$
func TestWordBreak(t *testing.T) {
	for _, tc := range []struct {
		s        string
		wordDict []string
		can      bool
	}{
		{s: "leetcode", wordDict: []string{"leet", "code"}, can: true},
		{s: "applepenapple", wordDict: []string{"apple", "pen"}, can: true},
		{s: "catsandog", wordDict: []string{"cats", "dog", "sand", "and", "cat"}, can: false},
	} {
		can := wordBreak(tc.s, tc.wordDict)
		assert.Equal(t, tc.can, can)
	}
}
