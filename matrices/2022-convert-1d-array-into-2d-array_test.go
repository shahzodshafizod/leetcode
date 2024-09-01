package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestConstruct2DArray$
func TestConstruct2DArray(t *testing.T) {
	for _, tc := range []struct {
		original []int
		m        int
		n        int
		matrix   [][]int
	}{
		{original: []int{1, 2, 3, 4}, m: 2, n: 2, matrix: [][]int{{1, 2}, {3, 4}}},
		{original: []int{1, 2, 3}, m: 1, n: 3, matrix: [][]int{{1, 2, 3}}},
		{original: []int{1, 2}, m: 1, n: 1, matrix: [][]int{}},
	} {
		matrix := construct2DArray(tc.original, tc.m, tc.n)
		assert.Equal(t, tc.matrix, matrix)
	}
}
