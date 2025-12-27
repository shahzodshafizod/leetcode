package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMatrixReshape$
func TestMatrixReshape(t *testing.T) {
	for _, tc := range []struct {
		mat    [][]int
		r      int
		c      int
		output [][]int
	}{
		{mat: [][]int{{1, 2}, {3, 4}}, r: 1, c: 4, output: [][]int{{1, 2, 3, 4}}},
		{mat: [][]int{{1, 2}, {3, 4}}, r: 2, c: 4, output: [][]int{{1, 2}, {3, 4}}},
		{mat: [][]int{{1, 2, 3, 4}}, r: 2, c: 2, output: [][]int{{1, 2}, {3, 4}}},
		{mat: [][]int{{1}}, r: 1, c: 1, output: [][]int{{1}}},
	} {
		output := matrixReshape(tc.mat, tc.r, tc.c)
		assert.Equal(t, tc.output, output)
	}
}
