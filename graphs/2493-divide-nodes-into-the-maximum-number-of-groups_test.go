package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMagnificentSets$
func TestMagnificentSets(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		groups int
	}{
		{n: 2, edges: [][]int{{1, 2}}, groups: 2},
		{n: 30, edges: [][]int{{16, 8}, {6, 5}}, groups: 30},
		{n: 3, edges: [][]int{{1, 2}, {2, 3}, {3, 1}}, groups: -1},
		{n: 6, edges: [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}, groups: 4},
		{n: 8, edges: [][]int{{5, 3}, {5, 8}, {7, 2}, {5, 4}, {6, 4}, {6, 3}, {5, 2}, {6, 2}, {1, 8}, {6, 8}, {7, 4}, {7, 3}, {7, 8}, {1, 4}}, groups: 4},
		{n: 24, edges: [][]int{{2, 13}, {7, 3}, {5, 3}, {21, 1}, {5, 1}, {4, 13}, {21, 19}, {7, 13}, {15, 3}, {21, 22}, {17, 19}, {23, 22}, {14, 13}}, groups: 19},
		{n: 14, edges: [][]int{{1, 4}, {3, 4}, {9, 8}, {12, 8}, {11, 8}, {11, 5}, {6, 4}, {3, 5}, {1, 8}, {13, 4}, {9, 5}, {9, 4}, {12, 5}, {7, 5}, {2, 4}, {3, 8}, {1, 5}, {12, 4}, {11, 4}, {10, 4}, {14, 5}, {14, 8}}, groups: 5},
	} {
		groups := magnificentSets(tc.n, tc.edges)
		assert.Equal(t, tc.groups, groups)
	}
}
