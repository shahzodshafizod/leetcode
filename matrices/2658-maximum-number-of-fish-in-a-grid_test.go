package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestFindMaxFish$
func TestFindMaxFish(t *testing.T) {
	for _, tc := range []struct {
		grid    [][]int
		maxfish int
	}{
		{grid: [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}, maxfish: 7},
		{grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}, maxfish: 1},
		{grid: [][]int{{0}}, maxfish: 0},
		{grid: [][]int{{4}}, maxfish: 4},
		{grid: [][]int{{0, 0}}, maxfish: 0},
		{grid: [][]int{{0, 6, 0, 9}}, maxfish: 9},
		{grid: [][]int{{4, 5, 5}, {0, 10, 0}}, maxfish: 24},
		{grid: [][]int{{3}, {4}, {0}, {4}, {8}, {9}, {0}}, maxfish: 21},
	} {
		maxfish := findMaxFish(tc.grid)
		assert.Equal(t, tc.maxfish, maxfish)
	}
}
