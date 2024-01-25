package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestWallsAndGates$
func TestWallsAndGates(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		result [][]int
	}{
		{matrix: [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}},
			result: [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}},
		{matrix: [][]int{{2147483647, -1, 0, 2147483647}, {-1, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}},
			result: [][]int{{2147483647, -1, 0, 1}, {-1, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}},
		{matrix: [][]int{}, result: [][]int{}},
		{matrix: [][]int{{}}, result: [][]int{{}}},
	} {
		result := wallsAndGates(tc.matrix)
		assert.Equal(t, tc.result, result)
	}
}
