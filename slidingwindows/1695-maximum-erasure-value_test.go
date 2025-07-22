package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMaximumUniqueSubarray$
func TestMaximumUniqueSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		score int
	}{
		{nums: []int{4, 2, 4, 5, 6}, score: 17},
		{nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}, score: 8},
	} {
		score := maximumUniqueSubarray(tc.nums)
		assert.Equal(t, tc.score, score)
	}
}
