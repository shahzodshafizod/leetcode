package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestLatestDayToCross$
func TestLatestDayToCross(t *testing.T) {
	for _, tc := range []struct {
		row   int
		col   int
		cells [][]int
		day   int
	}{
		{row: 2, col: 2, cells: [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}, day: 2},
		{row: 2, col: 2, cells: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, day: 1},
		{row: 3, col: 3, cells: [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}, day: 3},
	} {
		day := latestDayToCross(tc.row, tc.col, tc.cells)
		assert.Equal(t, tc.day, day)
	}
}
