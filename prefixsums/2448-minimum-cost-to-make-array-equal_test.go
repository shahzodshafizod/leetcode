package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestMinCost$
func TestMinCost(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		cost     []int
		minTotal int64
	}{
		{nums: []int{1, 3, 5, 2}, cost: []int{2, 3, 1, 14}, minTotal: 8},
		{nums: []int{2, 2, 2, 2, 2}, cost: []int{4, 2, 8, 1, 3}, minTotal: 0},
		// {nums: []int{735103, 366367, 132236, 133334, 808160, 113001, 49051, 735598, 686615, 665317, 999793, 426087, 587000, 649989, 509946, 743518}, cost: []int{724182, 447415, 723725, 902336, 600863, 287644, 13836, 665183, 448859, 917248, 397790, 898215, 790754, 320604, 468575, 825614}, minTotal: 1907611126748},
	} {
		minTotal := minCost(tc.nums, tc.cost)
		assert.Equal(t, tc.minTotal, minTotal)
	}
}
