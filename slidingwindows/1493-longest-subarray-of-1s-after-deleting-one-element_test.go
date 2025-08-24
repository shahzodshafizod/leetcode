package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestLongestSubarray1493$
func TestLongestSubarray1493(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		size int
	}{
		{nums: []int{1, 1, 0, 1}, size: 3},
		{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}, size: 5},
		{nums: []int{1, 1, 1}, size: 2},
	} {
		size := longestSubarray1493(tc.nums)
		assert.Equal(t, tc.size, size)
	}
}
