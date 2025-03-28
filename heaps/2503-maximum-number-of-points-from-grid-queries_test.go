package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestMaxPoints$
func TestMaxPoints(t *testing.T) {
	for _, tc := range []struct {
		grid    [][]int
		queries []int
		answer  []int
	}{
		{grid: [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, queries: []int{5, 6, 2}, answer: []int{5, 8, 1}},
		{grid: [][]int{{5, 2, 1}, {1, 1, 2}}, queries: []int{3}, answer: []int{0}},
	} {
		answer := maxPoints(tc.grid, tc.queries)
		assert.Equal(t, tc.answer, answer)
	}
}
