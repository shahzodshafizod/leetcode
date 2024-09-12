package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestCountConsistentStrings$
func TestCountConsistentStrings(t *testing.T) {
	for _, tc := range []struct {
		allowed string
		words   []string
		count   int
	}{
		{allowed: "ab", words: []string{"ad", "bd", "aaab", "baa", "badab"}, count: 2},
		{allowed: "abc", words: []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, count: 7},
		{allowed: "cad", words: []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, count: 4},
		{allowed: "xyz", words: []string{"x", "xy", "xyz", "z", "y", "xxy", "xxyzz"}, count: 7},
		{allowed: "a", words: []string{"a", "b", "c", "aa", "aaa", "ab", "ac"}, count: 3},
	} {
		count := countConsistentStrings(tc.allowed, tc.words)
		assert.Equal(t, tc.count, count)
	}
}
