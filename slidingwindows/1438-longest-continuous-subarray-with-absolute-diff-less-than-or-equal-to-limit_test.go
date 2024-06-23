package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestLongestSubarray$
func TestLongestSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		limit int
		size  int
	}{
		{nums: []int{8, 2, 4, 7}, limit: 4, size: 2},
		{nums: []int{10, 1, 2, 4, 7, 2}, limit: 5, size: 4},
		{nums: []int{4, 2, 2, 2, 4, 4, 2, 2}, limit: 0, size: 3},
	} {
		size := longestSubarray(tc.nums, tc.limit)
		assert.Equal(t, tc.size, size)
	}
}
