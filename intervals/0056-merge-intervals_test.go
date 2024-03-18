package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./intervals/ -run ^TestMerge$
func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		merged    [][]int
	}{
		{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, merged: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{intervals: [][]int{{1, 4}, {4, 5}}, merged: [][]int{{1, 5}}},
		{intervals: [][]int{{2, 5}, {1, 4}, {9, 10}, {6, 9}}, merged: [][]int{{1, 5}, {6, 10}}},
	} {
		merged := merge(tc.intervals)
		assert.Equal(t, tc.merged, merged)
	}
}
