package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMinimumSum$
func TestMinimumSum(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		area int
	}{
		{grid: [][]int{{1, 0, 1}, {1, 1, 1}}, area: 5},
		{grid: [][]int{{1, 0, 1, 0}, {0, 1, 0, 1}}, area: 5},
	} {
		area := minimumSum(tc.grid)
		assert.Equal(t, tc.area, area)
	}
}
