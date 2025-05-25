package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestLongestPalindrome$
func TestLongestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		words  []string
		length int
	}{
		{words: []string{"lc", "cl", "gg"}, length: 6},
		{words: []string{"cc", "ll", "xx"}, length: 2},
		{words: []string{"lc", "cl", "gg", "ak", "kk"}, length: 6},
		{words: []string{"ab", "ty", "yt", "lc", "cl", "ab"}, length: 8},
		{words: []string{"ab", "ty", "yt", "lc", "cl", "ab", "ba"}, length: 12},
		{words: []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}, length: 22},
		{words: []string{"qo", "fo", "fq", "qf", "fo", "ff", "qq", "qf", "of", "of", "oo", "of", "of", "qf", "qf", "of"}, length: 14},
	} {
		length := longestPalindrome(tc.words)
		assert.Equal(t, tc.length, length)
	}
}
