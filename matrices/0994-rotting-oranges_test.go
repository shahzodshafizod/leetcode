package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestOrangesRotting$
func TestOrangesRotting(t *testing.T) {
	for _, tc := range []struct {
		grid    [][]int
		minutes int
	}{
		{grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, minutes: 4},
		{grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, minutes: -1},
		{grid: [][]int{{0, 2}}, minutes: 0},
		{grid: [][]int{{2, 1, 1, 0, 0}, {1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}, minutes: 7},
		{grid: [][]int{{1, 1, 0, 0, 0}, {2, 1, 0, 0, 0}, {0, 0, 0, 1, 2}, {0, 1, 0, 0, 1}}, minutes: -1},
		{grid: [][]int{{2, 0, 1, 0, 0}, {1, 1, 0, 0, 2}, {0, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}, minutes: -1},
		{grid: [][]int{{0}}, minutes: 0},
	} {
		minutes := orangesRotting(tc.grid)
		assert.Equal(t, tc.minutes, minutes)
	}
}
