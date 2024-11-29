package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMinimumTime2577$
func TestMinimumTime2577(t *testing.T) {
	for _, tc := range []struct {
		grid    [][]int
		minTime int
	}{
		{grid: [][]int{{0, 2, 4}, {0, 5, 5}, {5, 4, 3}}, minTime: 6},
		{grid: [][]int{{0, 2, 4}, {3, 2, 1}, {1, 0, 4}}, minTime: -1},
		{grid: [][]int{{0, 5, 1}, {0, 7, 6}, {7, 7, 1}}, minTime: 8},
		{grid: [][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}, minTime: 7},
		{grid: [][]int{{0, 1, 12}, {19, 39, 97}, {75, 88, 33}, {21, 2, 88}}, minTime: 89},
		{grid: [][]int{{0, 7, 6, 6}, {1, 6, 8, 6}, {1, 5, 8, 3}, {4, 7, 0, 1}}, minTime: 10},
		// {grid: [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {25, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, minTime: 42},
	} {
		minTime := minimumTime2577(tc.grid)
		assert.Equal(t, tc.minTime, minTime)
	}
}
