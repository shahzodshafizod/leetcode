package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMinOperations2033$
func TestMinOperations2033(t *testing.T) {
	for _, tc := range []struct {
		grid  [][]int
		x     int
		opers int
	}{
		{grid: [][]int{{2, 4}, {6, 8}}, x: 2, opers: 4},
		{grid: [][]int{{1, 5}, {2, 3}}, x: 1, opers: 5},
		{grid: [][]int{{1, 2}, {3, 4}}, x: 2, opers: -1},
	} {
		opers := minOperations2033(tc.grid, tc.x)
		assert.Equal(t, tc.opers, opers)
	}
}
