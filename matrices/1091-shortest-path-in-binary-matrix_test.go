package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestShortestPathBinaryMatrix$
func TestShortestPathBinaryMatrix(t *testing.T) {
	for _, tc := range []struct {
		grid           [][]int
		shortestLength int
	}{
		{
			grid:           [][]int{{0, 1}, {1, 0}},
			shortestLength: 2,
		},
		{
			grid:           [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			shortestLength: 4,
		},
		{
			grid:           [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
			shortestLength: -1,
		},
		{
			grid:           [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}},
			shortestLength: -1,
		},
	} {
		shortestLength := shortestPathBinaryMatrix(tc.grid)
		assert.Equal(t, tc.shortestLength, shortestLength)
	}
}
