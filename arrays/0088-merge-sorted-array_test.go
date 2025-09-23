package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMerge$
func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		merged []int
	}{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, merged: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{1}, m: 1, nums2: []int{}, n: 0, merged: []int{1}},
		{nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, merged: []int{1}},
	} {
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		merged := tc.nums1
		assert.Equal(t, tc.merged, merged)
	}
}
