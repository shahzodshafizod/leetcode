package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestLengthOfLongestSubstring$
func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tc := range []struct {
		s         string
		maxLength int
	}{
		{s: "", maxLength: 0},
		{s: "abccabb", maxLength: 3},
		{s: "cccccc", maxLength: 1},
		{s: "a", maxLength: 1},
		{s: "abcbda", maxLength: 4},
		{s: "abcabcbb", maxLength: 3},
		{s: "bbbbb", maxLength: 1},
		{s: "pwwkew", maxLength: 3},
		{s: "abcbdca", maxLength: 4},
		{s: "abcbdaac", maxLength: 4},
		{s: "abccabb", maxLength: 3},
	} {
		maxLength := lengthOfLongestSubstring(tc.s)
		assert.Equal(t, tc.maxLength, maxLength)
	}
}
