package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindMedianSortedArrays$
func TestFindMedianSortedArrays(t *testing.T) {
	for _, tc := range []struct {
		nums1  []int
		nums2  []int
		median float64
	}{
		{nums1: []int{1, 2, 3, 4}, nums2: []int{5, 6, 7, 8}, median: 4.50000},
		{nums1: []int{1, 2}, nums2: []int{3, 4}, median: 2.50000},
		{nums1: []int{3, 8, 10}, nums2: []int{5, 9, 17}, median: 8.50000},
		{nums1: []int{1, 1}, nums2: []int{1, 2}, median: 1.00000},
		{nums1: []int{1, 3}, nums2: []int{2}, median: 2.00000},
		{nums1: []int{1}, nums2: []int{2, 3, 4, 5, 6}, median: 3.50000},
		{nums1: []int{1}, nums2: []int{2, 3, 4, 5, 6, 7}, median: 4},
		{nums1: []int{3, 4, 5, 6}, nums2: []int{1, 2, 3, 4}, median: 3.50000},
		{nums1: []int{3, 4, 5, 6, 7, 8}, nums2: []int{1, 2}, median: 4.50000},
		{nums1: []int{1, 2, 3, 4, 5, 6}, nums2: []int{7}, median: 4},
		{nums1: []int{1, 2, 3, 4, 5}, nums2: []int{6}, median: 3.50000},
		{nums1: []int{3, 4}, nums2: []int{1, 2}, median: 2.50000},
		{nums1: []int{5, 6, 7, 8}, nums2: []int{1, 2, 3, 4}, median: 4.50000},
		{nums1: []int{}, nums2: []int{1}, median: 1},
	} {
		median := findMedianSortedArrays(tc.nums1, tc.nums2)
		assert.InEpsilon(t, tc.median, median, 0.00001)
	}
}
