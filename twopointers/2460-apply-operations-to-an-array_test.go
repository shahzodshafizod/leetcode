package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestApplyOperations$
func TestApplyOperations(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		modified []int
	}{
		{nums: []int{1, 2, 2, 1, 1, 0}, modified: []int{1, 4, 2, 0, 0, 0}},
		{nums: []int{0, 1}, modified: []int{1, 0}},
		{nums: []int{1, 4, 0, 2, 0, 0}, modified: []int{1, 4, 2, 0, 0, 0}},
		{nums: []int{847, 847, 0, 0, 0, 399, 416, 416, 879, 879, 206, 206, 206, 272}, modified: []int{1694, 399, 832, 1758, 412, 206, 272, 0, 0, 0, 0, 0, 0, 0}},
	} {
		modified := applyOperations(tc.nums)
		assert.Equal(t, tc.modified, modified)
	}
}
