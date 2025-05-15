package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestGetLongestSubsequence$
func TestGetLongestSubsequence(t *testing.T) {
	for _, tc := range []struct {
		words       []string
		groups      []int
		alternating []string
	}{
		{words: []string{"e", "a", "b"}, groups: []int{0, 0, 1}, alternating: []string{"e", "b"}},
		{words: []string{"a", "b", "c", "d"}, groups: []int{1, 0, 1, 1}, alternating: []string{"a", "b", "c"}},
	} {
		alternating := getLongestSubsequence(tc.words, tc.groups)
		assert.Equal(t, tc.alternating, alternating)
	}
}
