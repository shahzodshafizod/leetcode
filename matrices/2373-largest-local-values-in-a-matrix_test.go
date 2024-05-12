package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestLargestLocal$
func TestLargestLocal(t *testing.T) {
	for _, tc := range []struct {
		grid     [][]int
		maxLocal [][]int
	}{
		{grid: [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}, maxLocal: [][]int{{9, 9}, {8, 6}}},
		{grid: [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}, maxLocal: [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}},
		{grid: [][]int{{9, 9, 8}, {5, 6, 2}, {8, 2, 6}}, maxLocal: [][]int{{9}}},
	} {
		maxLocal := largestLocal(tc.grid)
		assert.Equal(t, tc.maxLocal, maxLocal)
	}
}
