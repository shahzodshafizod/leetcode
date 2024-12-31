package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		ops  int
	}{
		{nums: []int{4, 2, 5, 3}, ops: 0},
		{nums: []int{1, 2, 3, 5, 6}, ops: 1},
		{nums: []int{1, 10, 100, 1000}, ops: 3},
	} {
		ops := minOperations(tc.nums)
		assert.Equal(t, tc.ops, ops)
	}
}
