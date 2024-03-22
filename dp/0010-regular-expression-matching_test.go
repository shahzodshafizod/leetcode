package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestIsMatch$
func TestIsMatch(t *testing.T) {
	for _, tc := range []struct {
		s       string
		p       string
		matches bool
	}{
		{s: "aa", p: "a", matches: false},
		{s: "aa", p: "a*", matches: true},
		{s: "ab", p: ".*", matches: true},
		{s: "aaa", p: "a*a", matches: true},
		{s: "mississippi", p: "mis*is*p*.", matches: false},
		{s: "mississippi", p: "mis*is*ip*.", matches: true},
		{s: "aab", p: "c*a*b", matches: true},
		{s: "abc", p: "a***abc", matches: true},
		{s: "aaa", p: "ab*ac*a", matches: true},
		{s: "ab", p: ".*c", matches: false},
		{s: "aaa", p: "aaaa", matches: false},
		{s: "aaa", p: "ab*a*c*a", matches: true},
		{s: "a", p: ".*..a*", matches: false},
	} {
		matches := isMatch(tc.s, tc.p)
		assert.Equal(t, tc.matches, matches)
	}
}
