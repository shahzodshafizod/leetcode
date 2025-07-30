package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestLongestSubarray$
func TestLongestSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		length int
	}{
		{nums: []int{1, 2, 3, 3, 2, 2}, length: 2},
		{nums: []int{1, 2, 3, 4}, length: 1},
	} {
		length := longestSubarray(tc.nums)
		assert.Equal(t, tc.length, length)
	}
}
