package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCherryPickup$
func TestCherryPickup(t *testing.T) {
	for _, tc := range []struct {
		grid  [][]int
		count int
	}{
		{grid: [][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}, count: 5},
		{grid: [][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}}, count: 0},
		{grid: [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 0, 1}, {1, 0, 1, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 1, 1, 1}}, count: 11},
		{grid: [][]int{{0, 1, 1, 0, 0}, {1, 1, 1, 1, 0}, {-1, 1, 1, 1, -1}, {0, 1, 1, 1, 0}, {1, 0, -1, 0, 0}}, count: 11},
		{grid: [][]int{{1, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 1}, {1, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 1, 1, 1}}, count: 15},
	} {
		count := cherryPickup(tc.grid)
		assert.Equal(t, tc.count, count)
	}
}
