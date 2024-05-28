package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestEqualSubstring$
func TestEqualSubstring(t *testing.T) {
	for _, tc := range []struct {
		s       string
		t       string
		maxCost int
		maxlen  int
	}{
		{s: "abcd", t: "bcdf", maxCost: 3, maxlen: 3},
		{s: "abcd", t: "cdef", maxCost: 3, maxlen: 1},
		{s: "abcd", t: "acde", maxCost: 0, maxlen: 1},
	} {
		maxlen := equalSubstring(tc.s, tc.t, tc.maxCost)
		assert.Equal(t, tc.maxlen, maxlen)
	}
}
