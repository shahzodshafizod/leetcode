package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestReverseVowels$
func TestReverseVowels(t *testing.T) {
	for _, tc := range []struct {
		s      string
		output string
	}{
		{s: "IceCreAm", output: "AceCreIm"},
		{s: "leetcode", output: "leotcede"},
		{s: "hello", output: "holle"},
		{s: "aA", output: "Aa"},
	} {
		output := reverseVowels(tc.s)
		assert.Equal(t, tc.output, output)
	}
}
