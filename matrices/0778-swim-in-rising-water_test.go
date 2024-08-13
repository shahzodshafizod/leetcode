package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestSwimInWater$
func TestSwimInWater(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		time int
	}{
		{grid: [][]int{{0, 2}, {1, 3}}, time: 3},
		{grid: [][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, time: 16},
	} {
		time := swimInWater(tc.grid)
		assert.Equal(t, tc.time, time)
	}
}
