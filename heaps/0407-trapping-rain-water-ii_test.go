package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestTrapRainWater$
func TestTrapRainWater(t *testing.T) {
	for _, tc := range []struct {
		heightMap [][]int
		volume    int
	}{
		{heightMap: [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}, volume: 4},
		{heightMap: [][]int{{3, 3, 3, 3, 3}, {3, 2, 2, 2, 3}, {3, 2, 1, 2, 3}, {3, 2, 2, 2, 3}, {3, 3, 3, 3, 3}}, volume: 10},
	} {
		volume := trapRainWater(tc.heightMap)
		assert.Equal(t, tc.volume, volume)
	}
}
