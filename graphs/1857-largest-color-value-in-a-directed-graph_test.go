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
		value  int
	}{
		{colors: "a", edges: [][]int{{0, 0}}, value: -1},
		{colors: "ab", edges: [][]int{{0, 1}, {1, 1}}, value: -1},
		{colors: "abaca", edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, value: 3},
		{colors: "aaaca", edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}, {4, 0}}, value: -1},
	} {
		value := largestPathValue(tc.colors, tc.edges)
		assert.Equal(t, tc.value, value)
	}
}
