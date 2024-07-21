package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestBuildMatrix$
func TestBuildMatrix(t *testing.T) {
	for _, tc := range []struct {
		k             int
		rowConditions [][]int
		colConditions [][]int
		matrix        [][]int
	}{
		{
			k:             3,
			rowConditions: [][]int{{1, 2}, {3, 2}},
			colConditions: [][]int{{2, 1}, {3, 2}},
			// matrix:        [][]int{{3, 0, 0}, {0, 0, 1}, {0, 2, 0}},
			matrix: [][]int{{0, 0, 1}, {3, 0, 0}, {0, 2, 0}},
		},
		{
			k:             3,
			rowConditions: [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}},
			colConditions: [][]int{{2, 1}},
			matrix:        nil, // [][]int{},
		},
	} {
		matrix := buildMatrix(tc.k, tc.rowConditions, tc.colConditions)
		assert.Equal(t, tc.matrix, matrix)
	}
}
