package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMaxCandies$
func TestMaxCandies(t *testing.T) {
	for _, tc := range []struct {
		status         []int
		candies        []int
		keys           [][]int
		containedBoxes [][]int
		initialBoxes   []int
		maxcount       int
	}{
		{status: []int{1, 1, 1}, candies: []int{100, 1, 100}, keys: [][]int{{}, {0, 2}, {}}, containedBoxes: [][]int{{}, {}, {}}, initialBoxes: []int{1}, maxcount: 1},
		{status: []int{1, 0, 1, 0}, candies: []int{7, 5, 4, 100}, keys: [][]int{{}, {}, {1}, {}}, containedBoxes: [][]int{{1, 2}, {3}, {}, {}}, initialBoxes: []int{0}, maxcount: 16},
		{status: []int{1, 0, 1, 0}, candies: []int{7, 5, 4, 100}, keys: [][]int{{}, {}, {1}, {3}}, containedBoxes: [][]int{{1, 2}, {3}, {}, {}}, initialBoxes: []int{0}, maxcount: 16},
		{status: []int{1, 0, 0, 0, 0, 0}, candies: []int{1, 1, 1, 1, 1, 1}, keys: [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}}, containedBoxes: [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}}, initialBoxes: []int{0}, maxcount: 6},
	} {
		maxcount := maxCandies(tc.status, tc.candies, tc.keys, tc.containedBoxes, tc.initialBoxes)
		assert.Equal(t, tc.maxcount, maxcount)
	}
}
