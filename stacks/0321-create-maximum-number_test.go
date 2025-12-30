package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestMaxNumber$
func TestMaxNumber(t *testing.T) {
	for _, tc := range []struct {
		nums1    []int
		nums2    []int
		k        int
		expected []int
	}{
		{[]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5, []int{9, 8, 6, 5, 3}},
		{[]int{6, 7}, []int{6, 0, 4}, 5, []int{6, 7, 6, 0, 4}},
		{[]int{3, 9}, []int{8, 9}, 3, []int{9, 8, 9}},
		{[]int{2, 5, 6, 4, 4, 0}, []int{7, 3, 8, 0, 6, 5, 7, 6, 2}, 15, []int{7, 3, 8, 2, 5, 6, 4, 4, 0, 6, 5, 7, 6, 2, 0}},
	} {
		output := maxNumber(tc.nums1, tc.nums2, tc.k)
		assert.Equal(t, tc.expected, output)
	}
}
