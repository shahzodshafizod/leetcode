package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestLargestIsland$
func TestLargestIsland(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		size int
	}{
		{grid: [][]int{{1, 0}, {0, 1}}, size: 3},
		{grid: [][]int{{1, 1}, {1, 0}}, size: 4},
		{grid: [][]int{{1, 1}, {1, 1}}, size: 4},
		{grid: [][]int{{0, 0}, {0, 0}}, size: 1},
		{grid: [][]int{{1, 1, 0, 1}, {1, 0, 0, 1}, {1, 0, 0, 1}, {1, 0, 0, 1}}, size: 10},
		{grid: [][]int{{1, 0, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}, {1, 0, 1, 0}}, size: 5},
		{grid: [][]int{{0, 0, 0, 1, 1}, {1, 0, 0, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 0, 1}, {0, 0, 0, 1, 0}}, size: 15},
	} {
		size := largestIsland(tc.grid)
		assert.Equal(t, tc.size, size)
	}
}
