package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestGetCommon$
func TestGetCommon(t *testing.T) {
	for _, tc := range []struct {
		nums1  []int
		nums2  []int
		common int
	}{
		{nums1: []int{1, 2, 3}, nums2: []int{2, 4}, common: 2},
		{nums1: []int{1, 2, 3, 6}, nums2: []int{2, 3, 4, 5}, common: 2},
		{nums1: []int{100, 200, 300}, nums2: []int{50, 100, 150, 200}, common: 100},
		{nums1: []int{2, 4, 6, 8, 10}, nums2: []int{1, 3, 5, 7, 9}, common: -1},
		{nums1: []int{2, 4, 6, 8, 10}, nums2: []int{1, 4, 5, 7, 9}, common: 4},
	} {
		common := getCommon(tc.nums1, tc.nums2)
		assert.Equal(t, tc.common, common)
	}
}
