package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMinimumOperations$
func TestMinimumOperations(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		output int
	}{
		{nums: []int{1, 2, 3, 4}, output: 3}, // 1->0, 2->3, 4->3: 3 operations
		{nums: []int{3, 6, 9}, output: 0},    // all divisible by 3
		{nums: []int{1, 2, 4, 5}, output: 4}, // all need 1 operation
		{nums: []int{0}, output: 0},          // 0 is divisible by 3
		{nums: []int{5}, output: 1},          // 5->6 or 5->3: 1 operation
	} {
		output := minimumOperations(tc.nums)
		assert.Equal(t, tc.output, output)
	}
}
