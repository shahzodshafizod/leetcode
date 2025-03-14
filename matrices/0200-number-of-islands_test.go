package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestNumIslands$
func TestNumIslands(t *testing.T) {
	for _, tc := range []struct {
		grid [][]byte
		num  int
	}{
		{grid: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, num: 1},
		{grid: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, num: 3},
		{grid: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '1'}, {'0', '0', '0', '1', '1'}}, num: 2},
		{grid: [][]byte{{'0', '1', '0', '1', '0'}, {'1', '0', '1', '0', '1'}, {'0', '1', '1', '1', '0'}, {'1', '0', '1', '0', '1'}}, num: 7},
		{grid: [][]byte{{'0', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, num: 0},
		{grid: [][]byte{{'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}}, num: 1},
	} {
		num := numIslands(tc.grid)
		assert.Equal(t, tc.num, num)
	}
}
