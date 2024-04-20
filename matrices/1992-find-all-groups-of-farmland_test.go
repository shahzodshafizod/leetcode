package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestFindFarmland$
func TestFindFarmland(t *testing.T) {
	for _, tc := range []struct {
		land     [][]int
		farmland [][]int
	}{
		{land: [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}, farmland: [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}}},
		{land: [][]int{{1, 1}, {1, 1}}, farmland: [][]int{{0, 0, 1, 1}}},
		{land: [][]int{{0}}, farmland: [][]int{}},
	} {
		farmland := findFarmland(tc.land)
		assert.Equal(t, tc.farmland, farmland)
	}
}
