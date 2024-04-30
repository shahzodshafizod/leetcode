package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestWonderfulSubstrings$
func TestWonderfulSubstrings(t *testing.T) {
	for _, tc := range []struct {
		word       string
		substrings int64
	}{
		{word: "aba", substrings: 4},
		{word: "aabb", substrings: 9},
		{word: "he", substrings: 2},
	} {
		substrings := wonderfulSubstrings(tc.word)
		assert.Equal(t, substrings, substrings)
	}
}
