package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestSumPrefixScores$
func TestSumPrefixScores(t *testing.T) {
	for _, tc := range []struct {
		words  []string
		answer []int
	}{
		{words: []string{"abc", "ab", "bc", "b"}, answer: []int{5, 4, 3, 2}},
		{words: []string{"abcd"}, answer: []int{4}},
		{words: []string{"a", "ab", "abc", "cab"}, answer: []int{3, 5, 6, 3}},
		{words: []string{"a"}, answer: []int{1}},
		{words: []string{"a", "b", "aa", "ab"}, answer: []int{3, 1, 4, 4}},
		{words: []string{"abababab", "abab", "ab", "a", "babababa", "baba", "ba", "b"}, answer: []int{15, 11, 7, 4, 15, 11, 7, 4}},
	} {
		answer := sumPrefixScores(tc.words)
		assert.Equal(t, tc.answer, answer)
	}
}
