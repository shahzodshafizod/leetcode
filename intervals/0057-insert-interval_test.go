package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./intervals/ -run ^TestInsert$
func TestInsert(t *testing.T) {
	for _, tc := range []struct {
		intervals   [][]int
		newInterval []int
		inserted    [][]int
	}{
		{intervals: [][]int{{1, 3}, {6, 9}}, newInterval: []int{2, 5}, inserted: [][]int{{1, 5}, {6, 9}}},
		{intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{4, 8}, inserted: [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{intervals: [][]int{{1, 3}, {8, 9}}, newInterval: []int{5, 7}, inserted: [][]int{{1, 3}, {5, 7}, {8, 9}}},
		{intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, newInterval: []int{6, 7}, inserted: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}},
		{intervals: [][]int{{4, 5}, {6, 9}}, newInterval: []int{1, 3}, inserted: [][]int{{1, 3}, {4, 5}, {6, 9}}},
	} {
		inserted := insert(tc.intervals, tc.newInterval)
		assert.Equal(t, tc.inserted, inserted)
	}
}
