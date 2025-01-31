package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestIsBipartite$
func TestIsBipartite(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		is    bool
	}{
		{graph: [][]int{{1}, {0}}, is: true},
		{graph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, is: false},
		{graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, is: true},
		{graph: [][]int{{4}, {}, {4}, {4}, {0, 2, 3}}, is: true},
		{graph: [][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}, is: false},
		{graph: [][]int{{3, 4, 6}, {3, 6}, {3, 6}, {0, 1, 2, 5}, {0, 7, 8}, {3}, {0, 1, 2, 7}, {4, 6}, {4}, {}}, is: true},
		{graph: [][]int{{1, 2, 3}, {0, 3, 4}, {0, 3}, {0, 1, 2}, {1}}, is: false},
		{graph: [][]int{{1, 2, 3, 4}, {0, 2, 3}, {0, 1, 3, 4}, {0, 1, 2, 4}, {0, 2, 3}}, is: false},
	} {
		is := isBipartite(tc.graph)
		assert.Equal(t, tc.is, is)
	}
}
