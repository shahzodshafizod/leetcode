package monotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/monotonic/ -run ^TestNextGreaterElement$
func TestNextGreaterElement(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		next  []int
	}{
		{nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}, next: []int{-1, 3, -1}},
		{nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}, next: []int{3, -1}},
	} {
		next := nextGreaterElement(tc.nums1, tc.nums2)
		assert.Equal(t, tc.next, next)
	}
}
