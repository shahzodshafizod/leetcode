package hashes

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestIntersect$
func TestIntersect(t *testing.T) {
	for _, tc := range []struct {
		nums1        []int
		nums2        []int
		intersection []int
	}{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, intersection: []int{2, 2}},
		{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, intersection: []int{4, 9}},
	} {
		intersection := intersect(tc.nums1, tc.nums2)
		sort.Ints(intersection)
		assert.Equal(t, tc.intersection, intersection)
	}
}
