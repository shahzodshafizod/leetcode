package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMinimumTime$
func TestMinimumTime(t *testing.T) {
	for _, tc := range []struct {
		n         int
		relations [][]int
		time      []int
		months    int
	}{
		{n: 1, relations: [][]int{}, time: []int{10}, months: 10},
		{n: 2, relations: [][]int{}, time: []int{3, 5}, months: 5},
		{n: 3, relations: [][]int{{1, 3}, {2, 3}}, time: []int{3, 2, 5}, months: 8},
		{n: 2, relations: [][]int{{2, 1}}, time: []int{10000, 10000}, months: 20000},
		{n: 10, relations: [][]int{{1, 3}}, time: []int{5, 95, 21, 41, 87, 9, 60, 67, 36, 96}, months: 96},
		{n: 5, relations: [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}, time: []int{1, 2, 3, 4, 5}, months: 12},
		{n: 7, relations: [][]int{{1, 7}, {1, 4}, {1, 3}, {2, 3}, {4, 3}, {5, 3}, {7, 3}, {7, 6}}, time: []int{6, 5, 1, 8, 2, 9, 4}, months: 19},
		{n: 7, relations: [][]int{{1, 7}, {1, 4}, {1, 3}, {2, 3}, {4, 3}, {5, 3}, {7, 3}, {7, 6}}, time: []int{6, 5, 1, 8, 2, 9, 4}, months: 19},
	} {
		months := minimumTime(tc.n, tc.relations, tc.time)
		assert.Equal(t, tc.months, months)
	}
}
