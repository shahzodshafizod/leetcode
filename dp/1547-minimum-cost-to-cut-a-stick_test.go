package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinCost$
func TestMinCost(t *testing.T) {
	for _, tc := range []struct {
		n    int
		cuts []int
		cost int
	}{
		{n: 7, cuts: []int{1, 3, 4, 5}, cost: 16},
		{n: 9, cuts: []int{5, 6, 1, 4, 2}, cost: 22},
	} {
		cost := minCost(tc.n, tc.cuts)
		assert.Equal(t, tc.cost, cost)
	}
}
