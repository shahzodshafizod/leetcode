package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSortedSquares$
func TestSortedSquares(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		sorted []int
	}{
		{nums: []int{-4, -1, 0, 3, 10}, sorted: []int{0, 1, 9, 16, 100}},
		{nums: []int{-7, -3, 2, 3, 11}, sorted: []int{4, 9, 9, 49, 121}},
		{nums: []int{-5, -3, -2, -1}, sorted: []int{1, 4, 9, 25}},
	} {
		sorted := sortedSquares(tc.nums)
		assert.Equal(t, tc.sorted, sorted)
	}
}
