package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCountCompleteComponents$
func TestCountCompleteComponents(t *testing.T) {
	for _, tc := range []struct {
		n     int
		edges [][]int
		count int
	}{
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}, count: 3},
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}, count: 1},
		{n: 5, edges: [][]int{{1, 2}, {3, 4}, {1, 4}, {2, 3}, {1, 3}, {2, 4}}, count: 2},
		{n: 3, edges: [][]int{{1, 0}, {2, 0}}, count: 0},
		{n: 4, edges: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, count: 0},
		{n: 2, edges: [][]int{{1, 0}}, count: 1},
	} {
		count := countCompleteComponents(tc.n, tc.edges)
		assert.Equal(t, tc.count, count)
	}
}
