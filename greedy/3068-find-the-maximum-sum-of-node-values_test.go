package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaximumValueSum$
func TestMaximumValueSum(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		edges [][]int
		sum   int64
	}{
		{nums: []int{1, 2, 1}, k: 3, edges: [][]int{{0, 1}, {0, 2}}, sum: 6},
		{nums: []int{2, 3}, k: 7, edges: [][]int{{0, 1}}, sum: 9},
		{nums: []int{7, 7, 7, 7, 7, 7}, k: 3, edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}, sum: 42},
		{nums: []int{24, 78, 1, 97, 44}, k: 6, edges: [][]int{{0, 2}, {1, 2}, {4, 2}, {3, 4}}, sum: 260},
	} {
		sum := maximumValueSum(tc.nums, tc.k, tc.edges)
		assert.Equal(t, tc.sum, sum)
	}
}
