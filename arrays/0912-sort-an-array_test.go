package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSortArray$
func TestSortArray(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		sorted []int
	}{
		{nums: []int{5, 2, 3, 1}, sorted: []int{1, 2, 3, 5}},
		{nums: []int{5, 1, 1, 2, 0, 0}, sorted: []int{0, 0, 1, 1, 2, 5}},
		{
			nums:   []int{1, 5, -7, 10, 5, -4, 125, 0, 2, 52, 14, 95, -1, 789, -32},
			sorted: []int{-32, -7, -4, -1, 0, 1, 2, 5, 5, 10, 14, 52, 95, 125, 789},
		},
		{
			nums:   []int{-32, -7, -4, -1, 0, 1, 2, 5, 5, 10, 14, 52, 95, 125, 789},
			sorted: []int{-32, -7, -4, -1, 0, 1, 2, 5, 5, 10, 14, 52, 95, 125, 789},
		},
		{
			nums:   []int{24, 7, 64, 53, 31, 28, -78, -9, 54, 28, -56, -70, 95, 43, 42, -64, 18, 13, 40, 81},
			sorted: []int{-78, -70, -64, -56, -9, 7, 13, 18, 24, 28, 28, 31, 40, 42, 43, 53, 54, 64, 81, 95},
		},
		{
			nums:   []int{1, 5, 10, 5, 0, 2, 52, 14, 95},
			sorted: []int{0, 1, 2, 5, 5, 10, 14, 52, 95},
		},
		{
			nums:   []int{24, 7, 64, 53, 31, 28, 54, 28, 82, 43, 42, 18, 13, 40, 81},
			sorted: []int{7, 13, 18, 24, 28, 28, 31, 40, 42, 43, 53, 54, 64, 81, 82},
		},
		{
			nums:   []int{6, 4, 7, 1, 4, 0, 3, 4, 5, 6, 2, 4, 2, 6, 11, 5, 45, 1, 6, 7, 5},
			sorted: []int{0, 1, 1, 2, 2, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 7, 7, 11, 45},
		},
	} {
		sorted := sortArray(tc.nums)
		assert.Equal(t, tc.sorted, sorted)
	}
}
