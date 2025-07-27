package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestMaxSubarrays$
func TestMaxSubarrays(t *testing.T) {
	for _, tc := range []struct {
		n                int
		conflictingPairs [][]int
		count            int64
	}{
		{n: 4, conflictingPairs: [][]int{{2, 3}, {1, 4}}, count: 9},
		{n: 5, conflictingPairs: [][]int{{1, 2}, {2, 5}, {3, 5}}, count: 12},
		{n: 10, conflictingPairs: [][]int{{1, 2}, {2, 3}, {3, 5}, {5, 7}, {7, 9}, {9, 10}, {4, 6}, {6, 8}, {8, 10}, {1, 10}}, count: 18},
		{n: 12, conflictingPairs: [][]int{{2, 4}, {5, 8}, {3, 9}, {1, 12}, {6, 11}, {4, 10}}, count: 51},
	} {
		count := maxSubarrays(tc.n, tc.conflictingPairs)
		assert.Equal(t, tc.count, count)
	}
}
