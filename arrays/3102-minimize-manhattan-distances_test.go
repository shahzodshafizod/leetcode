package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestMinimizeMaxDistance$
func TestMinimizeMaxDistance(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		output int
	}{
		{points: [][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}, output: 12},
		{points: [][]int{{1, 1}, {1, 1}, {1, 1}}, output: 0},
		{points: [][]int{{1, 2}, {3, 4}, {5, 6}}, output: 4},
	} {
		output := minimizeMaxDistance(tc.points)
		assert.Equal(t, tc.output, output)
	}
}
