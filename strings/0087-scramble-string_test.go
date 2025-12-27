package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestIsScramble$
func TestIsScramble(t *testing.T) {
	for _, tc := range []struct {
		s1     string
		s2     string
		output bool
	}{
		{s1: "great", s2: "rgeat", output: true},
		{s1: "abcde", s2: "caebd", output: false},
		{s1: "a", s2: "a", output: true},
		{s1: "abc", s2: "bca", output: true},
		{s1: "abcdbdacbdac", s2: "bdacabcdbdac", output: true},
		{s1: "abb", s2: "bab", output: true},
		{s1: "abcdefghijklmnopq", s2: "efghijklmnopqcadb", output: false},
	} {
		output := isScramble(tc.s1, tc.s2)
		assert.Equal(t, tc.output, output)
	}
}
