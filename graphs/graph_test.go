package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestTraversalBFS$
func TestTraversalBFS(t *testing.T) {
	for _, tc := range []struct {
		graph  [][]int
		values []int
	}{
		{
			graph:  [][]int{{1, 3}, {0}, {3, 8}, {4, 5, 2, 0}, {6, 3}, {3}, {7, 4}, {6}, {2}},
			values: []int{0, 1, 3, 4, 5, 2, 6, 8, 7},
		},
	} {
		values := traversalBFS(tc.graph)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./graphs/ -run ^TestTraversalDFS$
func TestTraversalDFS(t *testing.T) {
	for _, tc := range []struct {
		graph  [][]int
		values []int
	}{
		{
			graph:  [][]int{{1, 3}, {0}, {3, 8}, {4, 5, 2, 0}, {6, 3}, {3}, {7, 4}, {6}, {2}},
			values: []int{0, 1, 3, 4, 6, 7, 5, 2, 8},
		},
	} {
		values := traversalDFS(tc.graph)
		assert.Equal(t, tc.values, values)
	}
}
