package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestCountPaths2867$
func TestCountPaths2867(t *testing.T) {
	for _, tc := range []struct {
		n int
		edges [][]int
		count int64
	}{
		{n: 5, edges: [][]int{{1, 2}, {1, 3}, {2, 4}, {2, 5}}, count: 4},
		{n: 6, edges: [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 5}, {3, 6}}, count: 6},
		{n: 4, edges: [][]int{{1, 2}, {4, 1}, {3, 4}}, count: 4},
	} {
		count := countPaths2867(tc.n, tc.edges)
		assert.Equal(t, tc.count, count)
	}
}
