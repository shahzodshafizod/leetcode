package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestLargestPathValue$
func TestLargestPathValue(t *testing.T) {
	for _, tc := range []struct {
		colors string
		edges  [][]int
		color  int
	}{
		{colors: "a", edges: [][]int{{0, 0}}, color: -1},
		{colors: "ab", edges: [][]int{{0, 1}, {1, 1}}, color: -1},
		{colors: "abaca", edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, color: 3},
	} {
		color := largestPathValue(tc.colors, tc.edges)
		assert.Equal(t, tc.color, color)
	}
}
