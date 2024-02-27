package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxSubArray$
func TestMaxSubArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		max  int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, max: 6},
		{nums: []int{2, -1}, max: 2},
		{nums: []int{-2, -3, -1, -5}, max: -1},
		{nums: []int{1}, max: 1},
		{nums: []int{5, 4, -1, 7, 8}, max: 23},
		{nums: []int{5, 4, 7, 8}, max: 24},
		{nums: []int{1, 4, -1, 2}, max: 6},
	} {
		max := maxSubArray(tc.nums)
		assert.Equal(t, tc.max, max)
	}
}
