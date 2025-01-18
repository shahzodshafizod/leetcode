package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMinCost$
func TestMinCost(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		cost int
	}{
		{grid: [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}, cost: 3},
		{grid: [][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}, cost: 0},
		{grid: [][]int{{1, 2}, {4, 3}}, cost: 1},
		{grid: [][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}, cost: 3},
		{grid: [][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}, cost: 0},
		{grid: [][]int{{1, 1, 3}, {4, 3, 2}, {4, 2, 4}}, cost: 1},
	} {
		cost := minCost(tc.grid)
		assert.Equal(t, tc.cost, cost)
	}
}
