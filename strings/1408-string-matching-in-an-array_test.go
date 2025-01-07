package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestStringMatching$
func TestStringMatching(t *testing.T) {
	for _, tc := range []struct {
		words   []string
		matched []string
	}{
		{words: []string{"mass", "as", "hero", "superhero"}, matched: []string{"as", "hero"}},
		{words: []string{"leetcode", "et", "code"}, matched: []string{"et", "code"}},
		{words: []string{"blue", "green", "bu"}, matched: []string{}},
	} {
		matched := stringMatching(tc.words)
		assert.Equal(t, tc.matched, matched)
	}
}
