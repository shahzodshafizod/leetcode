package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestFindRedundantConnection$
func TestFindRedundantConnection(t *testing.T) {
	for _, tc := range []struct {
		edges [][]int
		edge  []int
	}{
		{edges: [][]int{{1, 2}, {1, 3}, {2, 3}}, edge: []int{2, 3}},
		{edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, edge: []int{1, 4}},
	} {
		edge := findRedundantConnection(tc.edges)
		assert.Equal(t, tc.edge, edge)
	}
}
