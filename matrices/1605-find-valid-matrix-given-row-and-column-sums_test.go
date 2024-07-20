package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestRestoreMatrix$
func TestRestoreMatrix(t *testing.T) {
	for _, tc := range []struct {
		rowSum []int
		colSum []int
		matrix [][]int
	}{
		{rowSum: []int{3, 8}, colSum: []int{4, 7}, matrix: [][]int{{3, 0}, {1, 7}}},
		{rowSum: []int{5, 7, 10}, colSum: []int{8, 6, 8}, matrix: [][]int{{5, 0, 0}, {3, 4, 0}, {0, 2, 8}}},
		// {rowSum: []int{5, 7, 10}, colSum: []int{8, 6, 8}, matrix: [][]int{{0, 5, 0}, {6, 1, 0}, {2, 0, 8}}},
	} {
		matrix := restoreMatrix(tc.rowSum, tc.colSum)
		assert.Equal(t, tc.matrix, matrix)
	}
}
