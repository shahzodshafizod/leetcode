package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestSortMatrix$
func TestSortMatrix(t *testing.T) {
	for _, tc := range []struct {
		grid   [][]int
		sorted [][]int
	}{
		{grid: [][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}, sorted: [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}}},
		{grid: [][]int{{0, 1}, {1, 2}}, sorted: [][]int{{2, 1}, {1, 0}}},
		{grid: [][]int{{1}}, sorted: [][]int{{1}}},
	} {
		sorted := sortMatrix(tc.grid)
		assert.Equal(t, tc.sorted, sorted)
	}
}
