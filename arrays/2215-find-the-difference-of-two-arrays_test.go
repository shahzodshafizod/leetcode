package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindDifference$
func TestFindDifference(t *testing.T) {
	for _, tc := range []struct {
		nums1  []int
		nums2  []int
		answer [][]int
	}{
		{nums1: []int{1, 2, 3}, nums2: []int{2, 4, 6}, answer: [][]int{{1, 3}, {4, 6}}},
		{nums1: []int{1, 2, 3, 3}, nums2: []int{1, 1, 2, 2}, answer: [][]int{{3}, {}}},
	} {
		answer := findDifference(tc.nums1, tc.nums2)
		assert.Equal(t, tc.answer, answer)
	}
}
