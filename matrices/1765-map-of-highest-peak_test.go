package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestHighestPeak$
func TestHighestPeak(t *testing.T) {
	for _, tc := range []struct {
		isWater [][]int
		height  [][]int
	}{
		{isWater: [][]int{{1}}, height: [][]int{{0}}},
		{isWater: [][]int{{1}, {0}}, height: [][]int{{0}, {1}}},
		{isWater: [][]int{{0, 1}, {0, 0}}, height: [][]int{{1, 0}, {2, 1}}},
		{isWater: [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}, height: [][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}}},
		{isWater: [][]int{{0}, {0}, {0}, {0}, {1}, {0}, {0}, {1}, {1}}, height: [][]int{{4}, {3}, {2}, {1}, {0}, {1}, {1}, {0}, {0}}},
	} {
		height := highestPeak(tc.isWater)
		assert.Equal(t, tc.height, height)
	}
}
