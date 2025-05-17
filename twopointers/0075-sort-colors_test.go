package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestSortColors$
func TestSortColors(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		sorted []int
	}{
		{nums: []int{2, 0, 2, 1, 1, 0}, sorted: []int{0, 0, 1, 1, 2, 2}},
		{nums: []int{2, 0, 1}, sorted: []int{0, 1, 2}},
	} {
		sortColors(tc.nums)
		assert.Equal(t, tc.sorted, tc.nums)
	}
}
