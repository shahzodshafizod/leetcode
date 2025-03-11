package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestNumberOfSubstrings$
func TestNumberOfSubstrings(t *testing.T) {
	for _, tc := range []struct {
		s     string
		count int
	}{
		{s: "abcabc", count: 10},
		{s: "aaacb", count: 3},
		{s: "abc", count: 1},
	} {
		count := numberOfSubstrings(tc.s)
		assert.Equal(t, tc.count, count)
	}
}
