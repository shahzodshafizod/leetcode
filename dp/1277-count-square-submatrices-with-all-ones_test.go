package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCountSquares$
func TestCountSquares(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		count  int
	}{
		{matrix: [][]int{{1, 0, 0}, {1, 1, 0}}, count: 3},
		{matrix: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, count: 7},
		{matrix: [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}, count: 15},
		{matrix: [][]int{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 1, 0, 0}, {1, 0, 0, 0}}, count: 9},
		{matrix: [][]int{{0, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 1}, {1, 0, 1, 0}}, count: 13},
		{matrix: [][]int{{1, 0, 1, 0, 1}, {1, 0, 0, 1, 1}, {0, 1, 0, 1, 1}, {1, 0, 0, 1, 1}}, count: 14},
		{matrix: [][]int{{1, 1, 0, 0, 1}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}, {0, 0, 1, 0, 1}}, count: 19},
	} {
		count := countSquares(tc.matrix)
		assert.Equal(t, tc.count, count)
	}
}
