package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestEventualSafeNodes$
func TestEventualSafeNodes(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		safe  []int
	}{
		{graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}, safe: []int{2, 4, 5, 6}},
		{graph: [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}, safe: []int{4}},
	} {
		safe := eventualSafeNodes(tc.graph)
		assert.Equal(t, tc.safe, safe)
	}
}
