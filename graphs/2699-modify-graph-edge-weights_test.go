package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestModifiedGraphEdges$
func TestModifiedGraphEdges(t *testing.T) {
	for _, tc := range []struct {
		n             int
		edges         [][]int
		source        int
		destination   int
		target        int
		modifiedEdges [][]int
	}{
		{n: 5, edges: [][]int{{4, 1, -1}, {2, 0, -1}, {0, 3, -1}, {4, 3, -1}}, source: 0, destination: 1, target: 5, modifiedEdges: [][]int{{4, 1, 1}, {2, 0, 3}, {0, 3, 3}, {4, 3, 1}}},
		{n: 3, edges: [][]int{{0, 1, -1}, {0, 2, 5}}, source: 0, destination: 2, target: 6, modifiedEdges: [][]int{}},
		{n: 4, edges: [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, -1}}, source: 0, destination: 2, target: 6, modifiedEdges: [][]int{{1, 0, 4}, {1, 2, 3}, {2, 3, 5}, {0, 3, 1}}},
		{n: 4, edges: [][]int{{0, 1, -1}, {1, 2, -1}, {3, 1, -1}, {3, 0, 2}, {0, 2, 5}}, source: 2, destination: 3, target: 8, modifiedEdges: [][]int{}},
		{n: 4, edges: [][]int{{0, 1, -1}, {2, 0, 2}, {3, 2, 6}, {2, 1, 10}, {3, 0, -1}}, source: 1, destination: 3, target: 12, modifiedEdges: [][]int{{0, 1, 11}, {2, 0, 2}, {3, 2, 6}, {2, 1, 10}, {3, 0, 1}}},
		{n: 5, edges: [][]int{{1, 4, 1}, {2, 4, -1}, {3, 0, 2}, {0, 4, -1}, {1, 3, 10}, {1, 0, 10}}, source: 0, destination: 2, target: 15, modifiedEdges: [][]int{{1, 4, 1}, {2, 4, 4}, {3, 0, 2}, {0, 4, 14}, {1, 3, 10}, {1, 0, 10}}},
	} {
		modifiedEdges := modifiedGraphEdges(tc.n, tc.edges, tc.source, tc.destination, tc.target)
		assert.Equal(t, tc.modifiedEdges, modifiedEdges)
	}
}
