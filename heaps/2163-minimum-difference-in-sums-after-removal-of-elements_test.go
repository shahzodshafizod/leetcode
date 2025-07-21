package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestMinimumDifference$
func TestMinimumDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		diff int64
	}{
		{nums: []int{3, 1, 2}, diff: -1},
		{nums: []int{7, 9, 5, 8, 1, 3}, diff: 1},
		{nums: []int{7, 9, 5, 8, 1, 3, 1, 1, 10000}, diff: -9993},
		{nums: []int{16, 46, 43, 41, 42, 14, 36, 49, 50, 28, 38, 25, 17, 5, 18, 11, 14, 21, 23, 39, 23}, diff: -14},
	} {
		diff := minimumDifference(tc.nums)
		assert.Equal(t, tc.diff, diff)
	}
}
