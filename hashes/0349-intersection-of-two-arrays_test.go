package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestIntersection$
func TestIntersection(t *testing.T) {
	for _, tc := range []struct {
		nums1  []int
		nums2  []int
		common []int
	}{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, common: []int{2}},
		{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, common: []int{9, 4}},
	} {
		common := intersection(tc.nums1, tc.nums2)
		assert.Equal(t, tc.common, common)
	}
}
