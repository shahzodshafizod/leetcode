package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMinDays$
func TestMinDays(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		days int
	}{
		{grid: [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}, days: 2},
		{grid: [][]int{{1, 1}}, days: 2},
		{grid: [][]int{{1, 0, 1, 0}}, days: 0},
	} {
		days := minDays(tc.grid)
		assert.Equal(t, tc.days, days)
	}
}
