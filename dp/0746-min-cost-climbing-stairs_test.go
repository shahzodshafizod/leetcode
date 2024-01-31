package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinCostClimbingStairs$
func TestMinCostClimbingStairs(t *testing.T) {
	for _, tc := range []struct {
		cost []int
		min  int
	}{
		{cost: []int{10, 15, 20}, min: 15},
		{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, min: 6},
	} {
		min := minCostClimbingStairs(tc.cost)
		assert.Equal(t, tc.min, min)
	}
}
