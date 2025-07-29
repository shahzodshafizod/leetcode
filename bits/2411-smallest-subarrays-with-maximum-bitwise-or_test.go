package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestSmallestSubarrays$
func TestSmallestSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		answer []int
	}{
		{nums: []int{1, 0, 2, 1, 3}, answer: []int{3, 3, 2, 2, 1}},
		{nums: []int{1, 2}, answer: []int{2, 1}},
	} {
		answer := smallestSubarrays(tc.nums)
		assert.Equal(t, tc.answer, answer)
	}
}
