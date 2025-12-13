package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestIntersectionSizeTwo$
func TestIntersectionSizeTwo(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		output    int
	}{
		{intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}, output: 3},
		{intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}, output: 5},
		{intervals: [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}, output: 8},
	} {
		output := intersectionSizeTwo(tc.intervals)
		assert.Equal(t, tc.output, output)
	}
}
