package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMaxAreaOfIsland$
func TestMaxAreaOfIsland(t *testing.T) {
	for _, tc := range []struct {
		grid    [][]int
		maxArea int
	}{
		{grid: [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, maxArea: 6},
		{grid: [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}, maxArea: 0},
		{grid: [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}, maxArea: 4},
	} {
		maxArea := maxAreaOfIsland(tc.grid)
		assert.Equal(t, tc.maxArea, maxArea)
	}
}
