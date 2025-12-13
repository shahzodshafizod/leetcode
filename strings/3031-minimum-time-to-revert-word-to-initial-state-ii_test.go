package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestMinimumTimeToInitialState$
func TestMinimumTimeToInitialState(t *testing.T) {
	for _, tc := range []struct {
		word   string
		k      int
		output int
	}{
		{word: "abacaba", k: 3, output: 2},
		{word: "abacaba", k: 4, output: 1},
		{word: "aabaa", k: 2, output: 2},
		{word: "abc", k: 1, output: 3},
	} {
		output := minimumTimeToInitialState(tc.word, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
