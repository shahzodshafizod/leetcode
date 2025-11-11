package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestFindMinimumTime$
func TestFindMinimumTime(t *testing.T) {
	for _, tc := range []struct {
		tasks  [][]int
		output int
	}{
		{[][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}, 2},
		{[][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}, 4},
		{[][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}, 2},
		{[][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}, 4},
		{[][]int{{1, 1, 1}}, 1},
		{[][]int{{1, 10, 1}, {2, 10, 1}, {3, 10, 1}}, 1}, // All use time 10
		{[][]int{{1, 5, 5}}, 5},                          // Must use all times 1,2,3,4,5
		{[][]int{{1, 4, 3}, {2, 4, 2}, {3, 4, 1}}, 3},    // Times 4,3,2 cover all
		{[][]int{{1, 5, 3}, {3, 6, 3}, {5, 9, 2}}, 4},
		{[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}}, 2}, // Times 2,3 can cover first two
	} {
		output := findMinimumTime(tc.tasks)
		assert.Equal(t, tc.output, output)
	}
}
