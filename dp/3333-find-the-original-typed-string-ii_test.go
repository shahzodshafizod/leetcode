package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestPossibleStringCount$
func TestPossibleStringCount(t *testing.T) {
	for _, tc := range []struct {
		word  string
		k     int
		count int
	}{
		{word: "aabbccdd", k: 7, count: 5},
		{word: "aabbccdd", k: 8, count: 1},
		{word: "aaabbb", k: 3, count: 8},
	} {
		count := possibleStringCount(tc.word, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
