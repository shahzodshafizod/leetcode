package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestSecondMinimum$
func TestSecondMinimum(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		time   int
		change int
		min    int
	}{
		{n: 5, edges: [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, time: 3, change: 5, min: 13},
		{n: 2, edges: [][]int{{1, 2}}, time: 3, change: 2, min: 11},
	} {
		minimum := secondMinimum(tc.n, tc.edges, tc.time, tc.change)
		assert.Equal(t, tc.min, minimum)
	}
}
