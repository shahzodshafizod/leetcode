package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinSum$
func TestMinSum(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		sum   int64
	}{
		{nums1: []int{3, 2, 0, 1, 0}, nums2: []int{6, 5, 0}, sum: 12},
		{nums1: []int{2, 0, 2, 0}, nums2: []int{1, 4}, sum: -1},
	} {
		sum := minSum(tc.nums1, tc.nums2)
		assert.Equal(t, tc.sum, sum)
	}
}
