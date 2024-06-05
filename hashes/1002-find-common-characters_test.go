package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCommonChars$
func TestCommonChars(t *testing.T) {
	for _, tc := range []struct {
		words  []string
		common []string
	}{
		{words: []string{"bella", "label", "roller"}, common: []string{"e", "l", "l"}},
		{words: []string{"cool", "lock", "cook"}, common: []string{"c", "o"}},
	} {
		common := commonChars(tc.words)
		assert.Equal(t, tc.common, common)
	}
}
