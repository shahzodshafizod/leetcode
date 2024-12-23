package monotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/monotonic/ -run ^TestLeftmostBuildingQueries$
func TestLeftmostBuildingQueries(t *testing.T) {
	for _, tc := range []struct {
		heights []int
		queries [][]int
		ans     []int
	}{
		{heights: []int{6, 4, 8, 5, 2, 7}, queries: [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}, ans: []int{2, 5, -1, 5, 2}},
		{heights: []int{5, 3, 8, 2, 6, 1, 4, 6}, queries: [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}}, ans: []int{7, 6, -1, 4, 6}},
	} {
		ans := leftmostBuildingQueries(tc.heights, tc.queries)
		assert.Equal(t, tc.ans, ans)
	}
}
