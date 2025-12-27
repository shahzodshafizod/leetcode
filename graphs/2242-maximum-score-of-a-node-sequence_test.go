package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMaximumScore$
func TestMaximumScore(t *testing.T) {
	for _, tc := range []struct {
		scores []int
		edges  [][]int
		output int
	}{
		{scores: []int{5, 2, 9, 8, 4}, edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}, output: 24},
		{scores: []int{9, 20, 6, 4, 11, 12}, edges: [][]int{{0, 3}, {5, 3}, {2, 4}, {1, 3}}, output: -1},
		{scores: []int{1, 1, 1, 1}, edges: [][]int{{0, 1}, {1, 2}, {2, 3}}, output: 4},
		{scores: []int{5, 10, 15, 20}, edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}}, output: 50},
	} {
		output := maximumScore(tc.scores, tc.edges)
		assert.Equal(t, tc.output, output)
	}
}
