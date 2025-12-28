package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestIsMatchWildcard$
func TestIsMatchWildcard(t *testing.T) {
	for _, tc := range []struct {
		s      string
		p      string
		result bool
	}{
		{s: "aa", p: "a", result: false},
		{s: "aa", p: "*", result: true},
		{s: "cb", p: "?a", result: false},
		{s: "adceb", p: "*a*b", result: true},
		{s: "acdcb", p: "a*c?b", result: false},
		{s: "", p: "*", result: true},
		{s: "", p: "?", result: false},
		{s: "a", p: "", result: false},
		{s: "", p: "", result: true},
		{s: "abc", p: "abc", result: true},
		{s: "abc", p: "a*c", result: true},
		{s: "abc", p: "a?c", result: true},
		{s: "abcdefg", p: "a*g", result: true},
		{s: "ho", p: "ho**", result: true},
		{s: "mississippi", p: "m*issi*ippi", result: true},
	} {
		result := isMatchWildcard(tc.s, tc.p)
		assert.Equal(t, tc.result, result)
	}
}
