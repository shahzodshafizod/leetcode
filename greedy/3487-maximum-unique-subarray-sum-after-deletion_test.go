package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxSum$
func TestMaxSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		sum  int
	}{
		{nums: []int{1, 2, 3, 4, 5}, sum: 15},
		{nums: []int{1, 1, 0, 1, 1}, sum: 1},
		{nums: []int{1, 2, -1, -2, 1, 0, -1}, sum: 3},
	} {
		sum := maxSum(tc.nums)
		assert.Equal(t, tc.sum, sum)
	}
}
