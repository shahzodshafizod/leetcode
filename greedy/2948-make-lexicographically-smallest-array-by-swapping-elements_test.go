package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestLexicographicallySmallestArray$
func TestLexicographicallySmallestArray(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		limit   int
		swapped []int
	}{
		{nums: []int{1, 5, 3, 9, 8}, limit: 2, swapped: []int{1, 3, 5, 8, 9}},
		{nums: []int{1, 7, 6, 18, 2, 1}, limit: 3, swapped: []int{1, 6, 7, 18, 1, 2}},
		{nums: []int{1, 7, 28, 19, 10}, limit: 3, swapped: []int{1, 7, 28, 19, 10}},
	} {
		swapped := lexicographicallySmallestArray(tc.nums, tc.limit)
		assert.Equal(t, tc.swapped, swapped)
	}
}
