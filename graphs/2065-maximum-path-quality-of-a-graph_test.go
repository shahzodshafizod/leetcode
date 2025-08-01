package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMaximalPathQuality$
func TestMaximalPathQuality(t *testing.T) {
	for _, tc := range []struct {
		values []int
		edges [][]int
		maxTime int
		quality int
	}{
		{values: []int{0, 32, 10, 43}, edges: [][]int{{0, 1, 10}, {1, 2, 15}, {0, 3, 10}}, maxTime: 49, quality: 75},
		{values: []int{5, 10, 15, 20}, edges: [][]int{{0, 1, 10}, {1, 2, 10}, {0, 3, 10}}, maxTime: 30, quality: 25},
		{values: []int{1, 2, 3, 4}, edges: [][]int{{0, 1, 10}, {1, 2, 11}, {2, 3, 12}, {1, 3, 13}}, maxTime: 50, quality: 7},
	} {
		quality := maximalPathQuality(tc.values, tc.edges, tc.maxTime)
		assert.Equal(t, tc.quality, quality)
	}
}
