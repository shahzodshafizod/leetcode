package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinCut$
func TestMinCut(t *testing.T) {
	for _, tc := range []struct {
		s      string
		result int
	}{
		{
			s:      "aab",
			result: 1,
		},
		{
			s:      "a",
			result: 0,
		},
		{
			s:      "ab",
			result: 1,
		},
		{
			s:      "ababababababababababababcbabababababababababababa",
			result: 0,
		},
		{
			s:      "abcdefg",
			result: 6,
		},
		{
			s:      "racecar",
			result: 0,
		},
		{
			s:      "aabbcc",
			result: 2,
		},
		{
			s:      "aaabaa",
			result: 1,
		},
	} {
		result := minCut(tc.s)
		assert.Equal(t, tc.result, result)
	}
}
