package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMinimumArea$
func TestMinimumArea(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		area int
	}{
		{grid: [][]int{{0, 1, 0}, {1, 0, 1}}, area: 6},
		{grid: [][]int{{1, 0}, {0, 0}}, area: 1},
	} {
		area := minimumArea(tc.grid)
		assert.Equal(t, tc.area, area)
	}
}
