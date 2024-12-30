package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestIsAlienSorted$
func TestIsAlienSorted(t *testing.T) {
	for _, tc := range []struct {
		words  []string
		order  string
		sorted bool
	}{
		{words: []string{"hello", "leetcode"}, order: "hlabcdefgijkmnopqrstuvwxyz", sorted: true},
		{words: []string{"word", "world", "row"}, order: "worldabcefghijkmnpqstuvxyz", sorted: false},
		{words: []string{"apple", "app"}, order: "abcdefghijklmnopqrstuvwxyz", sorted: false},
		{words: []string{"apple", "w"}, order: "abcdefghijklmnopqrstuvwxyz", sorted: true},
	} {
		sorted := isAlienSorted(tc.words, tc.order)
		assert.Equal(t, tc.sorted, sorted)
	}
}
