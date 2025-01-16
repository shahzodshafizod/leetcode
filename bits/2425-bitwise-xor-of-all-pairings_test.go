package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestXorAllNums$
func TestXorAllNums(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		num3  int
	}{
		{nums1: []int{2, 1, 3}, nums2: []int{10, 2, 5, 0}, num3: 13},
		{nums1: []int{1, 2}, nums2: []int{3, 4}, num3: 0},
		{nums1: []int{2}, nums2: []int{4}, num3: 6},
		{nums1: []int{27, 6, 15}, nums2: []int{8, 23, 15, 16, 25, 5}, num3: 28},
		{nums1: []int{10, 0, 15, 25, 29, 3, 2}, nums2: []int{12}, num3: 12},
		{nums1: []int{16, 15, 8, 6, 29, 2, 96, 29}, nums2: []int{24, 12, 13}, num3: 115},
		{nums1: []int{29, 10, 25, 6}, nums2: []int{24, 11, 1, 17, 11, 5, 1, 1, 27}, num3: 8},
	} {
		num3 := xorAllNums(tc.nums1, tc.nums2)
		assert.Equal(t, tc.num3, num3)
	}
}
