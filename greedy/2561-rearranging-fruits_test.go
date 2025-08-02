package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinCost$
func TestMinCost(t *testing.T) {
	for _, tc := range []struct {
		basket1 []int
		basket2 []int
		cost    int64
	}{
		{basket1: []int{4, 2, 2, 2}, basket2: []int{1, 4, 1, 2}, cost: 1},
		{basket1: []int{2, 3, 4, 1}, basket2: []int{3, 2, 5, 1}, cost: -1},
	} {
		cost := minCost(tc.basket1, tc.basket2)
		assert.Equal(t, tc.cost, cost)
	}
}
