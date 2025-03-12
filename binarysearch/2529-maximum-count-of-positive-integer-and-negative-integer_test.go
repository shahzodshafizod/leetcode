package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMaximumCount$
func TestMaximumCount(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{-2, -1, -1, 1, 2, 3}, count: 3},
		{nums: []int{-3, -2, -1, 0, 0, 1, 2}, count: 3},
		{nums: []int{5, 20, 66, 1314}, count: 4},
	} {
		count := maximumCount(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
