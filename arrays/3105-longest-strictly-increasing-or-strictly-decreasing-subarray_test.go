package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestLongestMonotonicSubarray$
func TestLongestMonotonicSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		length int
	}{
		{nums: []int{1, 4, 3, 3, 2}, length: 2},
		{nums: []int{3, 3, 3, 3}, length: 1},
		{nums: []int{3, 2, 1}, length: 3},
	} {
		length := longestMonotonicSubarray(tc.nums)
		assert.Equal(t, tc.length, length)
	}
}
