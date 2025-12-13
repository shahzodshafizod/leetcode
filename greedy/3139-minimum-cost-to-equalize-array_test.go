package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinCostToEqualizeArray$
func TestMinCostToEqualizeArray(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		cost1  int
		cost2  int
		output int
	}{
		{nums: []int{4, 1}, cost1: 5, cost2: 2, output: 15},
		{nums: []int{2, 3, 3, 3, 5}, cost1: 2, cost2: 1, output: 6},
		{nums: []int{3, 5, 3}, cost1: 1, cost2: 3, output: 4},
		{nums: []int{1, 14, 14, 15}, cost1: 2, cost2: 1, output: 20},
		{nums: []int{1, 1000000}, cost1: 1000000, cost2: 1000000, output: 998993007},
	} {
		output := minCostToEqualizeArray(tc.nums, tc.cost1, tc.cost2)
		assert.Equal(t, tc.output, output)
	}
}
