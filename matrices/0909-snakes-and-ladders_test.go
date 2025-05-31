package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestSnakesAndLadders$
func TestSnakesAndLadders(t *testing.T) {
	for _, tc := range []struct {
		board [][]int
		moves int
	}{
		{board: [][]int{{-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 35, -1, -1, 13, -1}, {-1, -1, -1, -1, -1, -1}, {-1, 15, -1, -1, -1, -1}}, moves: 4},
		{board: [][]int{{-1, -1}, {-1, 3}}, moves: 1},
	} {
		moves := snakesAndLadders(tc.board)
		assert.Equal(t, tc.moves, moves)
	}
}
