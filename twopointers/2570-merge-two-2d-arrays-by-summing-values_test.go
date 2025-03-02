package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestMergeArrays$
func TestMergeArrays(t *testing.T) {
	for _, tc := range []struct {
		nums1  [][]int
		nums2  [][]int
		merged [][]int
	}{
		{nums1: [][]int{{1, 2}, {2, 3}, {4, 5}}, nums2: [][]int{{1, 4}, {3, 2}, {4, 1}}, merged: [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}}},
		{nums1: [][]int{{2, 4}, {3, 6}, {5, 5}}, nums2: [][]int{{1, 3}, {4, 3}}, merged: [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}}},
		{nums1: [][]int{{1000, 1000}}, nums2: [][]int{{999, 999}}, merged: [][]int{{999, 999}, {1000, 1000}}},
		{nums1: [][]int{{1, 3}, {2, 2}, {3, 1}}, nums2: [][]int{{1, 1}, {2, 2}, {3, 3}}, merged: [][]int{{1, 4}, {2, 4}, {3, 4}}},
		{nums1: [][]int{{1, 555}, {3, 777}, {5, 999}}, nums2: [][]int{{2, 666}, {4, 888}}, merged: [][]int{{1, 555}, {2, 666}, {3, 777}, {4, 888}, {5, 999}}},
		{nums1: [][]int{{998, 100}, {999, 1000}}, nums2: [][]int{{997, 500}, {999, 1000}, {1000, 1000}}, merged: [][]int{{997, 500}, {998, 100}, {999, 2000}, {1000, 1000}}},
	} {
		merged := mergeArrays(tc.nums1, tc.nums2)
		assert.Equal(t, tc.merged, merged)
	}
}
