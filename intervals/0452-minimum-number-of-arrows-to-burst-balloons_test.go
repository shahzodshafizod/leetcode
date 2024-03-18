package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./intervals/ -run ^TestFindMinArrowShots$
func TestFindMinArrowShots(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		arrows int
	}{
		{points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, arrows: 2},
		{points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, arrows: 4},
		{points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, arrows: 2},
		{points: [][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}, arrows: 2},
		{points: [][]int{{-2147483646, -2147483645}, {2147483646, 2147483647}}, arrows: 2},
		{points: [][]int{{1, 2}}, arrows: 1},
		{points: [][]int{{2, 3}, {2, 3}}, arrows: 1},
		{points: [][]int{{77171087, 133597895}, {45117276, 135064454}, {80695788, 90089372}, {91705403, 110208054}, {52392754, 127005153}, {53999932, 118094992}, {11549676, 55543044}, {43947739, 128157751}, {55636226, 105334812}, {69348094, 125645633}}, arrows: 3},
	} {
		arrows := findMinArrowShots(tc.points)
		assert.Equal(t, tc.arrows, arrows)
	}
}
