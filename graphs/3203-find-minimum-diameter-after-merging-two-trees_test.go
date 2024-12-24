package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMinimumDiameterAfterMerge$
func TestMinimumDiameterAfterMerge(t *testing.T) {
	for _, tc := range []struct {
		edges1   [][]int
		edges2   [][]int
		diameter int
	}{
		{edges1: [][]int{}, edges2: [][]int{}, diameter: 1},
		{edges1: [][]int{}, edges2: [][]int{{0, 1}}, diameter: 2},
		{edges1: [][]int{{0, 1}, {0, 2}, {0, 3}}, edges2: [][]int{{0, 1}}, diameter: 3},
		{edges1: [][]int{}, edges2: [][]int{{0, 1}, {1, 2}}, diameter: 2},
		{edges1: [][]int{{0, 1}, {2, 0}, {3, 2}, {3, 6}, {8, 7}, {4, 8}, {5, 4}, {3, 5}, {3, 9}}, edges2: [][]int{{0, 1}, {0, 2}, {0, 3}}, diameter: 7},
		{edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}, edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 4}, {2, 5}, {3, 6}, {2, 7}}, diameter: 5},
	} {
		diameter := minimumDiameterAfterMerge(tc.edges1, tc.edges2)
		assert.Equal(t, tc.diameter, diameter)
	}
}
