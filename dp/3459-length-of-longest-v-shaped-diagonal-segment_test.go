package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestLenOfVDiagonal$
func TestLenOfVDiagonal(t *testing.T) {
	for _, tc := range []struct {
		grid   [][]int
		length int
	}{
		{grid: [][]int{{2, 2, 1, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}}, length: 5},
		{grid: [][]int{{2, 2, 2, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}}, length: 4},
		{grid: [][]int{{1, 2, 2, 2, 2}, {2, 2, 2, 2, 0}, {2, 0, 0, 0, 0}, {0, 0, 2, 2, 2}, {2, 0, 0, 2, 0}}, length: 5},
		{grid: [][]int{{1}}, length: 1},
	} {
		length := lenOfVDiagonal(tc.grid)
		assert.Equal(t, tc.length, length)
	}
}
