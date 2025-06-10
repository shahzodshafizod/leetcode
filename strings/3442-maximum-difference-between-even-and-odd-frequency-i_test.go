package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestMaxDifference$
func TestMaxDifference(t *testing.T) {
	for _, tc := range []struct {
		s    string
		diff int
	}{
		{s: "aaaaabbc", diff: 3},
		{s: "abcabcab", diff: 1},
	} {
		diff := maxDifference(tc.s)
		assert.Equal(t, tc.diff, diff)
	}
}
