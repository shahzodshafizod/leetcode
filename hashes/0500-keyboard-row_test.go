package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindWords$
func TestFindWords(t *testing.T) {
	for _, tc := range []struct {
		words []string
		res   []string
	}{
		{words: []string{"Hello", "Alaska", "Dad", "Peace"}, res: []string{"Alaska", "Dad"}},
		{words: []string{"omk"}, res: nil},
		{words: []string{"adsdf", "sfd"}, res: []string{"adsdf", "sfd"}},
	} {
		res := findWords(tc.words)
		assert.Equal(t, tc.res, res)
	}
}
