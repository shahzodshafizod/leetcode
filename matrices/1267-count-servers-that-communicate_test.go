package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestCountServers$
func TestCountServers(t *testing.T) {
	for _, tc := range []struct {
		grid  [][]int
		count int
	}{
		{grid: [][]int{{1, 0}, {0, 1}}, count: 0},
		{grid: [][]int{{1, 0}, {1, 1}}, count: 3},
		{grid: [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}, count: 4},
	} {
		count := countServers(tc.grid)
		assert.Equal(t, tc.count, count)
	}
}
