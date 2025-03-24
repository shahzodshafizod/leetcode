package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./intervals/ -run ^TestCountDays$
func TestCountDays(t *testing.T) {
	for _, tc := range []struct {
		days     int
		meetings [][]int
		count    int
	}{
		{days: 10, meetings: [][]int{{5, 7}, {1, 3}, {9, 10}}, count: 2},
		{days: 5, meetings: [][]int{{2, 4}, {1, 3}}, count: 1},
		{days: 6, meetings: [][]int{{1, 6}}, count: 0},
		{days: 14, meetings: [][]int{{6, 11}, {7, 13}, {8, 9}, {5, 8}, {3, 13}, {11, 13}, {1, 3}, {5, 10}, {8, 13}, {3, 9}}, count: 1},
	} {
		count := countDays(tc.days, tc.meetings)
		assert.Equal(t, tc.count, count)
	}
}
