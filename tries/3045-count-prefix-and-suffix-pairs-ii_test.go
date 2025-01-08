package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestCountPrefixSuffixPairs3045$
func TestCountPrefixSuffixPairs3045(t *testing.T) {
	for _, tc := range []struct {
		words []string
		count int64
	}{
		{words: []string{"a", "aba", "ababa", "aa"}, count: 4},
		{words: []string{"pa", "papa", "ma", "mama"}, count: 2},
		{words: []string{"abab", "ab"}, count: 0},
	} {
		count := countPrefixSuffixPairs3045(tc.words)
		assert.Equal(t, tc.count, count)
	}
}
