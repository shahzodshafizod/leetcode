package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestLargestRectangleArea$
func TestLargestRectangleArea(t *testing.T) {
	for _, tc := range []struct {
		heights []int
		largest int
	}{
		{heights: []int{2, 1, 5, 6, 2, 3}, largest: 10},
		{heights: []int{2, 1, 6, 5, 2, 3}, largest: 10},
		{heights: []int{2, 1, 2, 1, 5, 6, 2, 3, 1, 1, 1}, largest: 11},
		{heights: []int{1, 1, 2, 1, 5, 6, 2, 3, 1, 1, 1}, largest: 11},
		{heights: []int{1, 1, 2, 1, 5, 6, 2, 3, 1, 0, 1}, largest: 10},
		{heights: []int{2, 4}, largest: 4},
	} {
		largest := largestRectangleArea(tc.heights)
		assert.Equal(t, tc.largest, largest)
	}
}
