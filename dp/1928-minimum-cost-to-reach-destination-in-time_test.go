package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinCostToReachDestination$
func TestMinCostToReachDestination(t *testing.T) {
	for _, tc := range []struct {
		maxTime     int
		edges       [][]int
		passingFees []int
		output      int
	}{
		{
			maxTime:     30,
			edges:       [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}},
			passingFees: []int{5, 1, 2, 20, 20, 3},
			output:      11,
		},
		{
			maxTime:     29,
			edges:       [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}},
			passingFees: []int{5, 1, 2, 20, 20, 3},
			output:      48,
		},
		{
			maxTime:     25,
			edges:       [][]int{{0, 1, 10}, {1, 2, 10}, {2, 5, 10}, {0, 3, 1}, {3, 4, 10}, {4, 5, 15}},
			passingFees: []int{5, 1, 2, 20, 20, 3},
			output:      -1,
		},
	} {
		output := minCost1928(tc.maxTime, tc.edges, tc.passingFees)
		assert.Equal(t, tc.output, output)
	}
}
