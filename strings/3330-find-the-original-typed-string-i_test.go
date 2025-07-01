package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestPossibleStringCount$
func TestPossibleStringCount(t *testing.T) {
	for _, tc := range []struct {
		word  string
		count int
	}{
		{word: "abbcccc", count: 5},
		{word: "abcd", count: 1},
		{word: "aaaa", count: 4},
	} {
		count := possibleStringCount(tc.word)
		assert.Equal(t, tc.count, count)
	}
}
