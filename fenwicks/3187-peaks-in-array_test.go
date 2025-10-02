package fenwicks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./fenwicks/ -run ^TestCountOfPeaks$
func TestCountOfPeaks(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		counts  []int
	}{
		{nums: []int{3, 1, 4, 2, 5}, queries: [][]int{{2, 3, 4}, {1, 0, 4}}, counts: []int{0}},
		{nums: []int{4, 1, 4, 2, 1, 5}, queries: [][]int{{2, 2, 4}, {1, 0, 2}, {1, 0, 4}}, counts: []int{0, 1}},
		{nums: []int{4, 9, 4, 10, 7}, queries: [][]int{{2, 3, 2}, {2, 1, 3}, {1, 2, 3}}, counts: []int{0}},
		{nums: []int{5, 4, 8, 6}, queries: [][]int{{1, 2, 2}, {1, 1, 2}, {2, 1, 6}}, counts: []int{0, 0}},
	} {
		counts := countOfPeaks(tc.nums, tc.queries)
		assert.Equal(t, tc.counts, counts)
	}
}
