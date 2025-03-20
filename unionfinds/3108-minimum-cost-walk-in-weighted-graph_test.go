package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestMinimumCost$
func TestMinimumCost(t *testing.T) {
	for _, tc := range []struct {
		n      int
		edges  [][]int
		query  [][]int
		answer []int
	}{
		{n: 5, edges: [][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}}, query: [][]int{{0, 3}, {3, 4}}, answer: []int{1, -1}},
		{n: 3, edges: [][]int{{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}}, query: [][]int{{1, 2}}, answer: []int{0}},
		{n: 3, edges: [][]int{{1, 0, 4}, {0, 2, 5}, {0, 2, 3}, {0, 2, 14}, {0, 2, 12}, {2, 0, 14}, {0, 2, 4}}, query: [][]int{{2, 1}}, answer: []int{0}},
		{n: 6, edges: [][]int{{2, 5, 3}, {0, 3, 1}, {1, 4, 0}, {0, 3, 0}, {0, 2, 5}, {2, 0, 2}, {5, 1, 2}, {1, 3, 1}, {2, 1, 4}, {3, 2, 3}}, query: [][]int{{1, 5}, {0, 1}, {1, 5}, {3, 1}, {1, 2}}, answer: []int{0, 0, 0, 0, 0}},
		{n: 9, edges: [][]int{{1, 7, 7}, {5, 6, 2}, {3, 8, 5}, {3, 6, 3}}, query: [][]int{{3, 8}, {1, 4}, {5, 2}, {5, 2}, {1, 3}}, answer: []int{0, -1, -1, -1, -1}},
	} {
		answer := minimumCost(tc.n, tc.edges, tc.query)
		assert.Equal(t, tc.answer, answer)
	}
}
