package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestPivotArray$
func TestPivotArray(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		pivot    int
		modified []int
	}{
		{nums: []int{9, 12, 5, 10, 14, 3, 10}, pivot: 10, modified: []int{9, 5, 3, 10, 10, 12, 14}},
		{nums: []int{-3, 4, 3, 2}, pivot: 2, modified: []int{-3, 2, 4, 3}},
	} {
		modified := pivotArray(tc.nums, tc.pivot)
		assert.Equal(t, tc.modified, modified)
	}
}
