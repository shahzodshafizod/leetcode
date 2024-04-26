package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinFallingPathSum$
func TestMinFallingPathSum(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		sum  int
	}{
		{grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, sum: 13},
		{grid: [][]int{{7}}, sum: 7},
		{grid: [][]int{{-37, 51, -36, 34, -22}, {82, 4, 30, 14, 38}, {-68, -52, -92, 65, -85}, {-49, -3, -77, 8, -19}, {-60, -71, -21, -62, -73}}, sum: -268},
		{grid: [][]int{{1, 1, 2}, {2, 1, 1}, {1, 2, 3}}, sum: 3},
	} {
		sum := minFallingPathSum(tc.grid)
		assert.Equal(t, tc.sum, sum)
	}
}
