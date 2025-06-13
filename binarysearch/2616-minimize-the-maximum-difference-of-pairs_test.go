package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMinimizeMax$
func TestMinimizeMax(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		p    int
		diff int
	}{
		{nums: []int{10, 1, 2, 7, 1, 3}, p: 2, diff: 1},
		{nums: []int{4, 2, 1, 2}, p: 1, diff: 0},
		{nums: []int{8, 9, 1, 5, 4, 3, 6, 4, 3, 7}, p: 4, diff: 1},
	} {
		diff := minimizeMax(tc.nums, tc.p)
		assert.Equal(t, tc.diff, diff)
	}
}
