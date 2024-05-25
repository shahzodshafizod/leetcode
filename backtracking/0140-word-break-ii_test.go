package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestWordBreak$
func TestWordBreak(t *testing.T) {
	for _, tc := range []struct {
		s         string
		wordDict  []string
		sentences []string
	}{
		{
			s:         "catsanddog",
			wordDict:  []string{"cat", "cats", "and", "sand", "dog"},
			sentences: []string{"cat sand dog", "cats and dog"},
		},
		{
			s:         "pineapplepenapple",
			wordDict:  []string{"apple", "pen", "applepen", "pine", "pineapple"},
			sentences: []string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"},
		},
		{
			s:         "catsandog",
			wordDict:  []string{"cats", "dog", "sand", "and", "cat"},
			sentences: []string{},
		},
		{
			s:         "aaaaaaa",
			wordDict:  []string{"aaaa", "aa"},
			sentences: []string{},
		},
		{
			s:         "aaaaaaa",
			wordDict:  []string{"aaaa", "aaa"},
			sentences: []string{"aaaa aaa", "aaa aaaa"},
		},
	} {
		sentences := wordBreak(tc.s, tc.wordDict)
		assert.Equal(t, tc.sentences, sentences)
	}
}
