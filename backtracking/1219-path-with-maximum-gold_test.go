package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestGetMaximumGold$
func TestGetMaximumGold(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		max  int
	}{
		{grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}, max: 24},
		{grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}, max: 28},
	} {
		maximum := getMaximumGold(tc.grid)
		assert.Equal(t, tc.max, maximum)
	}
}
