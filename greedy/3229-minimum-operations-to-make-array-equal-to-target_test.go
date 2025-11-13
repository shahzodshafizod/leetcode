package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinimumOperations$
func TestMinimumOperations(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target []int
		ops    int64
	}{
		{nums: []int{3, 5, 1, 2}, target: []int{4, 6, 2, 4}, ops: 2},
		{nums: []int{1, 3, 2}, target: []int{2, 1, 4}, ops: 5},
		{nums: []int{3, 5, 1, 2}, target: []int{4, 6, 2, 4}, ops: 2},
		{nums: []int{1, 3, 2}, target: []int{2, 1, 4}, ops: 5},
		{nums: []int{1, 1, 1, 1}, target: []int{1, 1, 1, 1}, ops: 0},
		{nums: []int{5, 5, 5, 5}, target: []int{3, 3, 3, 3}, ops: 2},
		{nums: []int{1, 2, 3, 4, 5}, target: []int{5, 4, 3, 2, 1}, ops: 8},
		{nums: []int{10}, target: []int{20}, ops: 10},
		{nums: []int{100, 200, 300}, target: []int{150, 150, 150}, ops: 200},
		{nums: []int{1, 5, 10}, target: []int{10, 5, 1}, ops: 18},
	} {
		ops := minimumOperations(tc.nums, tc.target)
		assert.Equal(t, tc.ops, ops)
	}
}
