package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinimumReplacement$
func TestMinimumReplacement(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		ops  int64
	}{
		{nums: []int{3, 9, 3}, ops: 2},
		{nums: []int{1, 2, 3, 4, 5}, ops: 0},
		{nums: []int{19, 7, 2, 24, 11, 16, 1, 11, 23}, ops: 73},
	} {
		ops := minimumReplacement(tc.nums)
		assert.Equal(t, tc.ops, ops)
	}
}
