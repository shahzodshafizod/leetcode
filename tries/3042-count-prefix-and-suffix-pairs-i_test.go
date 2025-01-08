package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestCountPrefixSuffixPairs$
func TestCountPrefixSuffixPairs(t *testing.T) {
	for _, tc := range []struct {
		words []string
		count int
	}{
		{words: []string{"a", "aba", "ababa", "aa"}, count: 4},
		{words: []string{"pa", "papa", "ma", "mama"}, count: 2},
		{words: []string{"abab", "ab"}, count: 0},
	} {
		count := countPrefixSuffixPairs(tc.words)
		assert.Equal(t, tc.count, count)
	}
}
