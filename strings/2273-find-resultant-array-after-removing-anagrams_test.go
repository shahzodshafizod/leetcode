package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestRemoveAnagrams$
func TestRemoveAnagrams(t *testing.T) {
	for _, tc := range []struct {
		words []string
		res   []string
	}{
		{words: []string{"abba", "baba", "bbaa", "cd", "cd"}, res: []string{"abba", "cd"}},
		{words: []string{"a", "b", "c", "d", "e"}, res: []string{"a", "b", "c", "d", "e"}},
	} {
		res := removeAnagrams(tc.words)
		assert.Equal(t, tc.res, res)
	}
}
