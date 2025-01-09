package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestPrefixCount$
func TestPrefixCount(t *testing.T) {
	for _, tc := range []struct {
		words []string
		pref  string
		count int
	}{
		{words: []string{"pay", "attention", "practice", "attend"}, pref: "at", count: 2},
		{words: []string{"leetcode", "win", "loops", "success"}, pref: "code", count: 0},
	} {
		count := prefixCount(tc.words, tc.pref)
		assert.Equal(t, tc.count, count)
	}
}
