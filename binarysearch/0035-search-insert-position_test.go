package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestSearchInsert$
func TestSearchInsert(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		index  int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, index: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, index: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, index: 4},
	} {
		index := searchInsert(tc.nums, tc.target)
		assert.Equal(t, tc.index, index)
	}
}
