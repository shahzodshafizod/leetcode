package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestMaxNumEdgesToRemove$
func TestMaxNumEdgesToRemove(t *testing.T) {
	for _, tc := range []struct {
		n            int
		edges        [][]int
		removedEdges int
	}{
		{n: 4, edges: [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}, removedEdges: 2},
		{n: 4, edges: [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4}}, removedEdges: 0},
		{n: 4, edges: [][]int{{3, 2, 3}, {1, 1, 2}, {2, 3, 4}}, removedEdges: -1},
	} {
		removedEdges := maxNumEdgesToRemove(tc.n, tc.edges)
		assert.Equal(t, tc.removedEdges, removedEdges)
	}
}
