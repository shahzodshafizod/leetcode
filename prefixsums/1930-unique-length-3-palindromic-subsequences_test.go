package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestCountPalindromicSubsequence$
func TestCountPalindromicSubsequence(t *testing.T) {
	for _, tc := range []struct {
		s     string
		count int
	}{
		{s: "aabca", count: 3},
		{s: "adc", count: 0},
		{s: "bbcbaba", count: 4},
	} {
		count := countPalindromicSubsequence(tc.s)
		assert.Equal(t, tc.count, count)
	}
}
