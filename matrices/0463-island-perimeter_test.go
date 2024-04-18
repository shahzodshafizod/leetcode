package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestIslandPerimeter$
func TestIslandPerimeter(t *testing.T) {
	for _, tc := range []struct {
		grid      [][]int
		perimeter int
	}{
		{grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, perimeter: 16},
		{grid: [][]int{{1}}, perimeter: 4},
		{grid: [][]int{{1, 0}}, perimeter: 4},
		{grid: [][]int{{1, 1, 1}, {1, 0, 1}}, perimeter: 12},
		{grid: [][]int{{0, 1, 0}}, perimeter: 4},
		{grid: [][]int{{1}, {1}, {1}, {1}, {1}, {1}}, perimeter: 14},
		{grid: [][]int{{1, 1}, {1, 1}}, perimeter: 8},
	} {
		perimeter := islandPerimeter(tc.grid)
		assert.Equal(t, tc.perimeter, perimeter)
	}
}
