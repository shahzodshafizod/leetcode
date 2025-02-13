package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMaximumSum$
func TestMaximumSum(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		maxsum int
	}{
		{nums: []int{18, 43, 36, 13, 7}, maxsum: 54},
		{nums: []int{10, 12, 19, 14}, maxsum: -1},
		{nums: []int{4}, maxsum: -1},
		{nums: []int{5, 1, 6}, maxsum: -1},
		{nums: []int{4, 6, 10, 6}, maxsum: 12},
		{nums: []int{2, 1, 5, 5, 2, 4}, maxsum: 10},
	} {
		maxsum := maximumSum(tc.nums)
		assert.Equal(t, tc.maxsum, maxsum)
	}
}
