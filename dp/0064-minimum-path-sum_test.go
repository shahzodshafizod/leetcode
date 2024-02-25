package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinPathSum$
func TestMinPathSum(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		sum  int
	}{
		{grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, sum: 7},
		{grid: [][]int{{1, 2, 3}, {4, 5, 6}}, sum: 12},
		{grid: [][]int{{1, 2, 3}, {4, 5, 0}}, sum: 6},
		{grid: [][]int{{1, 2, 5}, {3, 2, 1}}, sum: 6},
		{grid: [][]int{{4}}, sum: 4},
		{grid: [][]int{{9, 9, 0, 8, 9, 0, 5, 7, 2, 2, 7, 0, 8, 0, 2, 4, 8}, {4, 4, 2, 7, 6, 0, 9, 7, 3, 2, 5, 4, 6, 5, 4, 8, 7}, {4, 9, 7, 0, 7, 9, 2, 4, 0, 2, 4, 4, 6, 2, 8, 0, 7}, {7, 7, 9, 6, 6, 4, 8, 4, 8, 7, 9, 4, 7, 6, 9, 6, 5}, {1, 3, 7, 5, 7, 9, 7, 3, 3, 3, 8, 3, 6, 5, 0, 3, 6}, {7, 1, 0, 7, 5, 0, 6, 6, 5, 3, 2, 6, 0, 0, 9, 5, 7}, {6, 5, 6, 3, 8, 1, 8, 6, 4, 4, 3, 4, 9, 9, 3, 3, 1}, {1, 0, 2, 9, 7, 9, 3, 1, 7, 5, 1, 8, 2, 8, 4, 7, 6}, {9, 6, 7, 7, 4, 1, 4, 0, 6, 5, 1, 9, 0, 3, 2, 1, 7}, {2, 0, 8, 7, 1, 7, 4, 3, 5, 6, 1, 9, 4, 0, 0, 2, 7}, {9, 8, 1, 3, 8, 7, 1, 2, 8, 3, 7, 3, 4, 6, 7, 6, 6}, {4, 8, 3, 8, 1, 0, 4, 4, 1, 0, 4, 1, 4, 4, 0, 3, 5}, {6, 3, 4, 7, 5, 4, 2, 2, 7, 9, 8, 4, 5, 6, 0, 3, 9}, {0, 4, 9, 7, 1, 0, 7, 7, 3, 2, 1, 4, 7, 6, 0, 0, 0}}, sum: 77},
	} {
		sum := minPathSum(tc.grid)
		assert.Equal(t, tc.sum, sum)
	}
}
