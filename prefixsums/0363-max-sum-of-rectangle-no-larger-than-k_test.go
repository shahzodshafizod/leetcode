package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestMaxSumSubmatrix$
func TestMaxSumSubmatrix(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		k      int
		output int
	}{
		{matrix: [][]int{{1, 0, 1}, {0, -2, 3}}, k: 2, output: 2},
		{matrix: [][]int{{2, 2, -1}}, k: 3, output: 3},
		{matrix: [][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, k: 8, output: 8},
		{matrix: [][]int{{2, 2, -1}}, k: 0, output: -1},
		{matrix: [][]int{{1}}, k: 1, output: 1},
		{matrix: [][]int{{-1}}, k: -1, output: -1},
	} {
		output := maxSumSubmatrix(tc.matrix, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
