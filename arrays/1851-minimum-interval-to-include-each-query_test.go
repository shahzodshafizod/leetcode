package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMinInterval$
func TestMinInterval(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		queries   []int
		answer    []int
	}{
		{intervals: [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, queries: []int{2, 3, 4, 5}, answer: []int{3, 3, 1, 4}},
		{intervals: [][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, queries: []int{2, 19, 5, 22}, answer: []int{2, -1, 4, 6}},
	} {
		answer := minInterval(tc.intervals, tc.queries)
		assert.Equal(t, tc.answer, answer)
	}
}
