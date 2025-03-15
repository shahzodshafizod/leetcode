package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestRob$
func TestRob(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		max  int
	}{
		{nums: []int{1, 2, 3, 1}, max: 4},
		{nums: []int{2, 7, 9, 3, 1}, max: 12},
		{nums: []int{2, 1, 1, 2}, max: 4},
		{nums: []int{2, 1, 7, 9}, max: 11},
		{nums: []int{100, 10, 1, 10, 100}, max: 201},
		{nums: []int{1, 3, 5, 9}, max: 12},
		{nums: []int{2, 4, 6, 8}, max: 12},
		{nums: []int{40, 2, 4, 10}, max: 50},
		{nums: []int{350}, max: 350},
		// {nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, max: 0},
	} {
		max := rob(tc.nums)
		assert.Equal(t, tc.max, max)
	}
}
