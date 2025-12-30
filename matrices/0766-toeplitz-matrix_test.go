package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestIsToeplitzMatrix$
func TestIsToeplitzMatrix(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		output bool
	}{
		{matrix: [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, output: true},
		{matrix: [][]int{{1, 2}, {2, 2}}, output: false},
		{matrix: [][]int{{1}}, output: true},
		{matrix: [][]int{{1, 2}, {3, 1}}, output: true}, // Diagonals: [3], [1,1], [2]
		{matrix: [][]int{{11, 74, 0, 93}, {40, 11, 74, 0}, {13, 40, 11, 74}}, output: true},
	} {
		output := isToeplitzMatrix(tc.matrix)
		assert.Equal(t, tc.output, output)
	}
}
