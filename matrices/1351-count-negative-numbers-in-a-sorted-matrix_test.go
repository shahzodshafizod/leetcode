package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestCountNegatives$
func TestCountNegatives(t *testing.T) {
	for _, tc := range []struct {
		grid   [][]int
		output int
	}{
		{
			grid:   [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			output: 8,
		},
		{grid: [][]int{{3, 2}, {1, 0}}, output: 0},
		{grid: [][]int{{1, -1}, {-1, -1}}, output: 3},
		{grid: [][]int{{-1}}, output: 1},
		{grid: [][]int{{5, 1, 0}, {-5, -5, -5}}, output: 3},
		{
			grid:   [][]int{{3, -1, -3, -3, -3}, {2, -2, -3, -3, -3}, {1, -2, -3, -3, -3}, {0, -3, -3, -3, -3}},
			output: 16,
		},
	} {
		output := countNegatives(tc.grid)
		assert.Equal(t, tc.output, output)
	}
}
