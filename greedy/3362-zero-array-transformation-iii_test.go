package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxRemoval$
func TestMaxRemoval(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		maxRem  int
	}{
		{nums: []int{1, 2, 3, 4}, queries: [][]int{{0, 3}}, maxRem: -1},
		{nums: []int{2, 0, 2}, queries: [][]int{{0, 2}, {0, 2}, {1, 1}}, maxRem: 1},
		{nums: []int{0, 0, 1, 1, 0}, queries: [][]int{{3, 4}, {0, 2}, {2, 3}}, maxRem: 2},
		{nums: []int{0, 0, 1, 1, 0, 0}, queries: [][]int{{2, 3}, {0, 2}, {3, 5}}, maxRem: 2},
		{nums: []int{0, 0, 3}, queries: [][]int{{0, 2}, {1, 1}, {0, 0}, {0, 0}}, maxRem: -1},
		{nums: []int{1, 1, 1, 1}, queries: [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}, maxRem: 2},
		{nums: []int{1, 1, 1, 1}, queries: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 3}}, maxRem: 3},
		{nums: []int{0, 3}, queries: [][]int{{0, 1}, {0, 0}, {0, 1}, {0, 1}, {0, 0}}, maxRem: 2},
		{nums: []int{1, 2}, queries: [][]int{{1, 1}, {0, 0}, {1, 1}, {1, 1}, {0, 1}, {0, 0}}, maxRem: 4},
		{nums: []int{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0}, queries: [][]int{{0, 6}, {7, 14}, {6, 7}}, maxRem: 2},
	} {
		maxRem := maxRemoval(tc.nums, tc.queries)
		assert.Equal(t, tc.maxRem, maxRem)
	}
}
