package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target []int
		output int
	}{
		{nums: []int{1, 2, 3}, target: []int{4}, output: 1},
		{nums: []int{8, 4}, target: []int{10, 5}, output: 2},
		{nums: []int{7, 9, 10}, target: []int{7}, output: 0},
		{nums: []int{1, 2, 8}, target: []int{5, 10}, output: 2},
		{nums: []int{3, 5, 7}, target: []int{6}, output: 1},
	} {
		output := minOperations(tc.nums, tc.target)
		assert.Equal(t, tc.output, output)
	}
}
