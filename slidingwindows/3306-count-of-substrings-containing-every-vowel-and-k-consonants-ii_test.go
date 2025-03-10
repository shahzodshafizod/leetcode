package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestCountOfSubstrings$
func TestCountOfSubstrings(t *testing.T) {
	for _, tc := range []struct {
		word  string
		k     int
		count int64
	}{
		{word: "aeioqq", k: 1, count: 0},
		{word: "aeiou", k: 0, count: 1},
		{word: "ieaouqqieaouqq", k: 1, count: 3},
		{word: "iqeaouqi", k: 2, count: 3},
		{word: "aeouih", k: 0, count: 1},
		{word: "aadieuoh", k: 1, count: 2},
		{word: "euaoei", k: 1, count: 0},
		{word: "aoaiuefi", k: 1, count: 4},
	} {
		count := countOfSubstrings(tc.word, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
