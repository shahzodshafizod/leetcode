package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestLongestPalindrome$
func TestLongestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		subs string
	}{
		{s: "babad", subs: "bab"},
		{s: "dabab", subs: "aba"},
		{s: "cbbd", subs: "bb"},
		{s: "racecar", subs: "racecar"},
		{s: "rakuta", subs: "r"},
		{s: "forgeeksskeegfor", subs: "geeksskeeg"},
		{s: "a", subs: "a"},
		{s: "aa", subs: "aa"},
		{s: "aba", subs: "aba"},
		{s: "abba", subs: "abba"},
	} {
		subs := longestPalindrome(tc.s)
		assert.Equal(t, tc.subs, subs)
	}
}
