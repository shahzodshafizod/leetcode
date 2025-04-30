package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestRepeatedSubstringPattern$
func TestRepeatedSubstringPattern(t *testing.T) {
	for _, tc := range []struct {
		s     string
		canbe bool
	}{
		{s: "a", canbe: false},
		{s: "aa", canbe: true},
		{s: "aba", canbe: false},
		{s: "abab", canbe: true},
		{s: "abac", canbe: false},
		{s: "abaababaab", canbe: true},
		{s: "abcabcabcabc", canbe: true},
		{s: "babbabbabbabbab", canbe: true},
		{s: "ababababababaababababababaababababababa", canbe: true},
	} {
		canbe := repeatedSubstringPattern(tc.s)
		assert.Equal(t, tc.canbe, canbe)
	}
}
