package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestFindWordsContaining$
func TestFindWordsContaining(t *testing.T) {
	for _, tc := range []struct {
		words   []string
		x       byte
		indices []int
	}{
		{words: []string{"leet", "code"}, x: 'e', indices: []int{0, 1}},
		{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: 'a', indices: []int{0, 2}},
		{words: []string{"abc", "bcd", "aaaa", "cbc"}, x: 'z', indices: nil},
	} {
		indices := findWordsContaining(tc.words, tc.x)
		assert.Equal(t, tc.indices, indices)
	}
}
