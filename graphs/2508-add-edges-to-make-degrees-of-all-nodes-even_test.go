package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestIsPossible$
func TestIsPossible(t *testing.T) {
	for _, tc := range []struct {
		n        int
		edges    [][]int
		possible bool
	}{
		{n: 5, edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 4}, {2, 5}}, possible: true},
		{n: 4, edges: [][]int{{1, 2}, {3, 4}}, possible: true},
		{n: 4, edges: [][]int{{1, 2}, {1, 3}, {1, 4}}, possible: false},
	} {
		possible := isPossible(tc.n, tc.edges)
		assert.Equal(t, tc.possible, possible)
	}
}
