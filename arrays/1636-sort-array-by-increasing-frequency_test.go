package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFrequencySort$
func TestFrequencySort(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		sorted []int
	}{
		{nums: []int{1, 1, 2, 2, 2, 3}, sorted: []int{3, 1, 1, 2, 2, 2}},
		{nums: []int{2, 3, 1, 3, 2}, sorted: []int{1, 3, 3, 2, 2}},
		{nums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, sorted: []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
		{nums: []int{1, 5, 0, 5}, sorted: []int{1, 0, 5, 5}},
	} {
		sorted := frequencySort(tc.nums)
		assert.Equal(t, tc.sorted, sorted)
	}
}
