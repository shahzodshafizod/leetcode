package fenwicks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./fenwicks/ -run ^TestGoodTriplets$
func TestGoodTriplets(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		count int64
	}{
		{nums1: []int{2, 0, 1, 3}, nums2: []int{0, 1, 2, 3}, count: 1},
		{nums1: []int{4, 0, 1, 3, 2}, nums2: []int{4, 1, 0, 2, 3}, count: 4},
	} {
		count := goodTriplets(tc.nums1, tc.nums2)
		assert.Equal(t, tc.count, count)
	}
}
