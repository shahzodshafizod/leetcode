package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestMaximizeSumOfWeights$
func TestMaximizeSumOfWeights(t *testing.T) {
	for _, tc := range []struct {
		edges   [][]int
		k       int
		weights int64
	}{
		{edges: [][]int{{0, 1, 4}, {0, 2, 2}, {2, 3, 12}, {2, 4, 6}}, k: 2, weights: 22},
		{edges: [][]int{{0, 1, 5}, {1, 2, 10}, {0, 3, 15}, {3, 4, 20}, {3, 5, 5}, {0, 6, 10}}, k: 3, weights: 65},
		{edges: [][]int{{0, 1, 34}, {0, 2, 17}}, k: 1, weights: 34},
		{edges: [][]int{{0, 2, 10}, {0, 1, 23}}, k: 1, weights: 23},
	} {
		weights := maximizeSumOfWeights(tc.edges, tc.k)
		assert.Equal(t, tc.weights, weights)
	}
}
