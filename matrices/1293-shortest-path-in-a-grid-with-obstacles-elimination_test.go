package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestShortestPath$
func TestShortestPath(t *testing.T) {
	for _, tc := range []struct {
		grid   [][]int
		k      int
		length int
	}{
		{grid: [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}, k: 1, length: 6},
		{grid: [][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}, k: 1, length: -1},
	} {
		length := shortestPath(tc.grid, tc.k)
		assert.Equal(t, tc.length, length)
	}
}
