package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestLongestPalindrome$
func TestLongestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		s      string
		maxlen int
	}{
		{s: "abccccdd", maxlen: 7},
		{s: "a", maxlen: 1},
		{s: "aa", maxlen: 2},
	} {
		maxlen := longestPalindrome(tc.s)
		assert.Equal(t, tc.maxlen, maxlen)
	}
}
