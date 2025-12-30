package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestFindMin$
func TestFindMin(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output int
	}{
		{[]int{1, 3, 5}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
		{[]int{3, 3, 1, 3}, 1},
		{[]int{1, 1}, 1},
		{[]int{2, 2, 2, 2, 2}, 2},
		{[]int{3, 1, 3}, 1},
		{[]int{10, 1, 10, 10, 10}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 0},
		{[]int{7, 0, 1, 2, 4, 5, 6}, 0},
	} {
		output := findMin(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
