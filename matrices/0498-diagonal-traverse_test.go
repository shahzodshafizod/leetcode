package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestFindDiagonalOrder$
func TestFindDiagonalOrder(t *testing.T) {
	for _, tc := range []struct {
		mat    [][]int
		result []int
	}{
		{mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, result: []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		{mat: [][]int{{1, 2}, {3, 4}}, result: []int{1, 2, 3, 4}},
	} {
		result := findDiagonalOrder(tc.mat)
		assert.Equal(t, tc.result, result)
	}
}
