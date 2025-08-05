package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMaxTotalFruits$
func TestMaxTotalFruits(t *testing.T) {
	for _, tc := range []struct {
		fruits   [][]int
		startPos int
		k        int
		total    int
	}{
		{fruits: [][]int{{2, 8}, {6, 3}, {8, 6}}, startPos: 5, k: 4, total: 9},
		{fruits: [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, startPos: 5, k: 4, total: 14},
		{fruits: [][]int{{0, 3}, {6, 4}, {8, 5}}, startPos: 3, k: 2, total: 0},
		{fruits: [][]int{{200000, 10000}}, startPos: 200000, k: 0, total: 10000},
		{fruits: [][]int{{0, 7}, {7, 4}, {9, 10}, {12, 6}, {14, 8}, {16, 5}, {17, 8}, {19, 4}, {20, 1}, {21, 3}, {24, 3}, {25, 3}, {26, 1}, {28, 10}, {30, 9}, {31, 6}, {32, 1}, {37, 5}, {40, 9}}, startPos: 21, k: 30, total: 71},
	} {
		total := maxTotalFruits(tc.fruits, tc.startPos, tc.k)
		assert.Equal(t, tc.total, total)
	}
}
