package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMaximumImportance$
func TestMaximumImportance(t *testing.T) {
	for _, tc := range []struct {
		n          int
		roads      [][]int
		importance int64
	}{
		{n: 5, roads: [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}, importance: 43},
		{n: 5, roads: [][]int{{0, 3}, {2, 4}, {1, 3}}, importance: 20},
		{n: 5, roads: [][]int{{0, 1}}, importance: 9},
	} {
		importance := maximumImportance(tc.n, tc.roads)
		assert.Equal(t, tc.importance, importance)
	}
}
