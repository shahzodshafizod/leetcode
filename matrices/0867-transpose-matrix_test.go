package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestTranspose$
func TestTranspose(t *testing.T) {
	for _, tc := range []struct {
		matrix     [][]int
		transposed [][]int
	}{
		{matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, transposed: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{matrix: [][]int{{1, 2, 3}, {4, 5, 6}}, transposed: [][]int{{1, 4}, {2, 5}, {3, 6}}},
		{
			matrix:     [][]int{{10, 6, 3, 5}, {10, 2, 4, 8}, {3, 8, 3, 10}, {7, 8, 8, 6}},
			transposed: [][]int{{10, 10, 3, 7}, {6, 2, 8, 8}, {3, 4, 3, 8}, {5, 8, 10, 6}},
		},
	} {
		transposed := transpose(tc.matrix)
		assert.Equal(t, tc.transposed, transposed)
	}
}
