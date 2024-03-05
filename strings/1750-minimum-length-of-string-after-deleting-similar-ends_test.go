package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestMinimumLength$
func TestMinimumLength(t *testing.T) {
	for _, tc := range []struct {
		s      string
		length int
	}{
		{s: "ca", length: 2},
		{s: "cabaabac", length: 0},
		{s: "aabccabba", length: 3},
		{s: "bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb", length: 1},
		{s: "aaaaaaaaaaaaa", length: 0},
		{s: "aaaaaabbbbbababababababbabaabababbccccccbbbbbbaaa", length: 29},
		{s: "a", length: 1},
	} {
		length := minimumLength(tc.s)
		assert.Equal(t, tc.length, length)
	}
}
