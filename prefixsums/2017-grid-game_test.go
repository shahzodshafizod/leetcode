package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestGridGame$
func TestGridGame(t *testing.T) {
	for _, tc := range []struct {
		grid   [][]int
		robot2 int64
	}{
		{grid: [][]int{{2, 5, 4}, {1, 5, 1}}, robot2: 4},
		{grid: [][]int{{3, 3, 1}, {8, 5, 2}}, robot2: 4},
		{grid: [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}, robot2: 7},
	} {
		robot2 := gridGame(tc.grid)
		assert.Equal(t, tc.robot2, robot2)
	}
}

