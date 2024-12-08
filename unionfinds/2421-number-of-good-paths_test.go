package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestNumberOfGoodPaths$
func TestNumberOfGoodPaths(t *testing.T) {
	for _, tc := range []struct {
		vals  []int
		edges [][]int
		count int
	}{
		{vals: []int{1}, edges: [][]int{}, count: 1},
		{vals: []int{2, 1, 1}, edges: [][]int{{0, 1}, {2, 0}}, count: 3},
		{vals: []int{1, 3, 2, 1, 3}, edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, count: 6},
		{vals: []int{1, 1, 2, 2, 3}, edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}, count: 7},
	} {
		count := numberOfGoodPaths(tc.vals, tc.edges)
		assert.Equal(t, tc.count, count)
	}
}
