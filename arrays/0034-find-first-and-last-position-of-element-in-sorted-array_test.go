package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSearchRange$
func TestSearchRange(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		target    int
		positions []int
	}{
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 8, positions: []int{3, 4}},
		{nums: []int{5, 7, 7, 8, 8, 10}, target: 6, positions: []int{-1, -1}},
		{nums: []int{1, 3, 3, 5, 5, 5, 8, 9}, target: 5, positions: []int{3, 5}},
		{nums: []int{1, 2, 3, 4, 5, 6}, target: 4, positions: []int{3, 3}},
		{nums: []int{1, 2, 3, 4, 5}, target: 9, positions: []int{-1, -1}},
		{nums: []int{}, target: 3, positions: []int{-1, -1}},
		{nums: []int{4, 4, 4, 4, 4, 4}, target: 4, positions: []int{0, 5}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 13, 13, 13, 14, 15, 16}, target: 13, positions: []int{12, 15}},
	} {
		positions := searchRange(tc.nums, tc.target)
		assert.Equal(t, tc.positions, positions)
	}
}
