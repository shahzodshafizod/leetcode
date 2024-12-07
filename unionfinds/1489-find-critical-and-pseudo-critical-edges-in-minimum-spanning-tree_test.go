package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestFindCriticalAndPseudoCriticalEdges$
func TestFindCriticalAndPseudoCriticalEdges(t *testing.T) {
	for _, tc := range []struct {
		n       int
		edges   [][]int
		indices [][]int
	}{
		{n: 2, edges: [][]int{{0, 1, 3}}, indices: [][]int{{0}, {}}},
		{n: 3, edges: [][]int{{0, 1, 1}, {0, 2, 2}, {1, 2, 3}}, indices: [][]int{{0, 1}, {}}},
		{n: 4, edges: [][]int{{0, 1, 1}, {0, 2, 2}, {0, 3, 3}}, indices: [][]int{{0, 1, 2}, {}}},
		{n: 4, edges: [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}, indices: [][]int{{}, {0, 1, 2, 3}}},
		{n: 5, edges: [][]int{{0, 1, 2}, {0, 2, 2}, {2, 3, 2}, {2, 4, 2}}, indices: [][]int{{0, 1, 2, 3}, {}}},
		{n: 6, edges: [][]int{{0, 1, 2}, {0, 2, 5}, {2, 3, 5}, {1, 4, 4}, {2, 5, 5}, {4, 5, 2}}, indices: [][]int{{0, 5, 3, 2}, {1, 4}}},
		{n: 4, edges: [][]int{{0, 1, 1}, {0, 3, 1}, {0, 2, 1}, {1, 2, 1}, {1, 3, 1}, {2, 3, 1}}, indices: [][]int{{}, {0, 1, 2, 3, 4, 5}}},
		{n: 5, edges: [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}, indices: [][]int{{0, 1}, {2, 3, 4, 5}}},
		{n: 6, edges: [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 1}, {2, 3, 4}, {3, 4, 2}, {3, 5, 2}, {4, 5, 2}}, indices: [][]int{{3}, {0, 1, 2, 4, 5, 6}}},
		{n: 8, edges: [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 4, 1}, {1, 5, 1}, {2, 6, 1}, {3, 7, 1}, {4, 5, 1}, {6, 7, 1}}, indices: [][]int{{1}, {0, 2, 3, 4, 5, 6, 7, 8}}},
		{n: 6, edges: [][]int{{0, 1, 4}, {0, 2, 3}, {1, 3, 4}, {3, 4, 1}, {3, 5, 2}, {0, 5, 1}, {2, 3, 5}, {2, 4, 6}, {0, 4, 4}, {2, 5, 2}}, indices: [][]int{{3, 5, 4, 9}, {0, 2}}},
	} {
		indices := findCriticalAndPseudoCriticalEdges(tc.n, tc.edges)
		assert.Equal(t, tc.indices, indices)
	}
}
