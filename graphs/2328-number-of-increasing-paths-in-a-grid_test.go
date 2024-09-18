package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCountPaths$
func TestCountPaths(t *testing.T) {
	for _, tc := range []struct {
		grid  [][]int
		count int
	}{
		{grid: [][]int{{1, 1}, {3, 4}}, count: 8},
		{grid: [][]int{{1}, {2}}, count: 3},
	} {
		count := countPaths(tc.grid)
		assert.Equal(t, tc.count, count)
	}
}
