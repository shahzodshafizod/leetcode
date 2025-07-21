package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestRunningSum$
func TestRunningSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		rsum []int
	}{
		{nums: []int{1, 2, 3, 4}, rsum: []int{1, 3, 6, 10}},
		{nums: []int{1, 1, 1, 1, 1}, rsum: []int{1, 2, 3, 4, 5}},
		{nums: []int{3, 1, 2, 10, 1}, rsum: []int{3, 4, 6, 16, 17}},
	} {
		rsum := runningSum(tc.nums)
		assert.Equal(t, tc.rsum, rsum)
	}
}
