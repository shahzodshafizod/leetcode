package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMinimumScore$
func TestMinimumScore(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		edges [][]int
		score int
	}{
		{nums: []int{1, 5, 5, 4, 11}, edges: [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, score: 9},
		{nums: []int{5, 5, 2, 4, 4, 2}, edges: [][]int{{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3}}, score: 0},
	} {
		score := minimumScore(tc.nums, tc.edges)
		assert.Equal(t, tc.score, score)
	}
}
