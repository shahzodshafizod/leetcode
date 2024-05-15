package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestMaximumSafenessFactor$
func TestMaximumSafenessFactor(t *testing.T) {
	for _, tc := range []struct {
		grid   [][]int
		factor int
	}{
		{grid: [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}, factor: 0},
		{grid: [][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}, factor: 2},
		{grid: [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}, factor: 2},
		{grid: [][]int{{1}}, factor: 0},
		{grid: [][]int{{0, 1, 1}, {0, 0, 0}, {1, 1, 0}}, factor: 1},
		{grid: [][]int{{0, 1, 1}, {0, 0, 0}, {0, 1, 0}}, factor: 1},
	} {
		factor := maximumSafenessFactor(tc.grid)
		assert.Equal(t, tc.factor, factor)
	}
}
