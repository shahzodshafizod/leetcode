package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestPlacedCoins$
func TestPlacedCoins(t *testing.T) {
	for _, tc := range []struct {
		edges [][]int
		cost  []int
		coin  []int64
	}{
		{edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}, cost: []int{1, 2, 3, 4, 5, 6}, coin: []int64{120, 1, 1, 1, 1, 1}},
		{edges: [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 6}, {2, 7}, {2, 8}}, cost: []int{1, 4, 2, 3, 5, 7, 8, -4, 2}, coin: []int64{280, 140, 32, 1, 1, 1, 1, 1, 1}},
		{edges: [][]int{{0, 1}, {0, 2}}, cost: []int{1, 2, -2}, coin: []int64{0, 1, 1}},
	} {
		coin := placedCoins(tc.edges, tc.cost)
		assert.Equal(t, tc.coin, coin)
	}
}
