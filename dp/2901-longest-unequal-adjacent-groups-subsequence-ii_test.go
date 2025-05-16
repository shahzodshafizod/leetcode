package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestGetWordsInLongestSubsequence$
func TestGetWordsInLongestSubsequence(t *testing.T) {
	for _, tc := range []struct {
		words       []string
		groups      []int
		subsequence []string
	}{
		{words: []string{"bab", "dab", "cab"}, groups: []int{1, 2, 2}, subsequence: []string{"bab", "dab"}},
		{words: []string{"a", "b", "c", "d"}, groups: []int{1, 2, 3, 4}, subsequence: []string{"a", "b", "c", "d"}},
		{words: []string{"bdb", "aaa", "ada"}, groups: []int{2, 1, 3}, subsequence: []string{"aaa", "ada"}},
		// {words: []string{"bad", "dc", "bc", "ccd", "dd", "da", "cad", "dba", "aba"}, groups: []int{9, 7, 1, 2, 6, 8, 3, 7, 2}, subsequence: []string{"dc", "dd", "da"}},
	} {
		subsequence := getWordsInLongestSubsequence(tc.words, tc.groups)
		assert.Equal(t, tc.subsequence, subsequence)
	}
}
