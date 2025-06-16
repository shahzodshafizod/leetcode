package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMaximumDifference$
func TestMaximumDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		diff int
	}{
		{nums: []int{7, 1, 5, 4}, diff: 4},
		{nums: []int{9, 4, 3, 2}, diff: -1},
		{nums: []int{1, 5, 2, 10}, diff: 9},
	} {
		diff := maximumDifference(tc.nums)
		assert.Equal(t, tc.diff, diff)
	}
}
