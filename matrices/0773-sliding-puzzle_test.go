package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestSlidingPuzzle$
func TestSlidingPuzzle(t *testing.T) {
	for _, tc := range []struct {
		board [][]int
		moves int
	}{
		{board: [][]int{{1, 2, 3}, {4, 0, 5}}, moves: 1},
		{board: [][]int{{1, 2, 3}, {5, 4, 0}}, moves: -1},
		{board: [][]int{{4, 1, 2}, {5, 0, 3}}, moves: 5},
		{board: [][]int{{1, 3, 4}, {0, 2, 5}}, moves: 14},
		{board: [][]int{{3, 2, 4}, {1, 5, 0}}, moves: 14},
		{board: [][]int{{2, 4, 1}, {5, 3, 0}}, moves: 12},
		{board: [][]int{{0, 5, 2}, {4, 3, 1}}, moves: 15},
		{board: [][]int{{4, 2, 0}, {5, 1, 3}}, moves: 7},
		{board: [][]int{{3, 0, 1}, {4, 2, 5}}, moves: -1},
		{board: [][]int{{3, 2, 1}, {4, 0, 5}}, moves: -1},
		{board: [][]int{{1, 2, 3}, {4, 5, 0}}, moves: 0},
	} {
		moves := slidingPuzzle(tc.board)
		assert.Equal(t, tc.moves, moves)
	}
}
