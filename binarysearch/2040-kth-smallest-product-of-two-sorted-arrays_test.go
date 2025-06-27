package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestKthSmallestProduct$
func TestKthSmallestProduct(t *testing.T) {
	for _, tc := range []struct {
		nums1   []int
		nums2   []int
		k       int64
		product int64
	}{
		{nums1: []int{2, 5}, nums2: []int{3, 4}, k: 2, product: 8},
		{nums1: []int{-4, -2, 0, 3}, nums2: []int{2, 4}, k: 6, product: 0},
		{nums1: []int{-2, -1, 0, 1, 2}, nums2: []int{-3, -1, 2, 4, 5}, k: 3, product: -6},
		{nums1: []int{-6}, nums2: []int{-9}, k: 1, product: 54},
	} {
		product := kthSmallestProduct(tc.nums1, tc.nums2, tc.k)
		assert.Equal(t, tc.product, product)
	}
}
