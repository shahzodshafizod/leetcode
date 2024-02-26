package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestLongestPalindromeSubseq$
func TestLongestPalindromeSubseq(t *testing.T) {
	for _, tc := range []struct {
		s      string
		length int
	}{
		{s: "bbbab", length: 4},
		{s: "cbbd", length: 2},
		{s: "bbbb", length: 4},
		{s: "bbabb", length: 5},
		{s: "a", length: 1},
	} {
		length := longestPalindromeSubseq(tc.s)
		assert.Equal(t, tc.length, length)
	}
}
