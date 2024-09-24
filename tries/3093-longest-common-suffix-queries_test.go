package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestStringIndices$
func TestStringIndices(t *testing.T) {
	for _, tc := range []struct {
		wordsContainer []string
		wordsQuery     []string
		ans            []int
	}{
		{wordsContainer: []string{"abcd", "bcd", "xbcd"}, wordsQuery: []string{"cd", "bcd", "xyz"}, ans: []int{1, 1, 1}},
		{wordsContainer: []string{"abcdefgh", "poiuygh", "ghghgh"}, wordsQuery: []string{"gh", "acbfgh", "acbfegh"}, ans: []int{2, 0, 2}},
		{wordsContainer: []string{"abcde", "abcde"}, wordsQuery: []string{"abcde", "bcde", "cde", "de", "e"}, ans: []int{0, 0, 0, 0, 0}},
		{wordsContainer: []string{"abcd", "bcda"}, wordsQuery: []string{"cdf", "afa"}, ans: []int{0, 1}},
	} {
		ans := stringIndices(tc.wordsContainer, tc.wordsQuery)
		assert.Equal(t, tc.ans, ans)
	}
}
