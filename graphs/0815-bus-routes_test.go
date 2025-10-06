package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestNumBusesToDestination$
func TestNumBusesToDestination(t *testing.T) {
	for _, tc := range []struct {
		routes [][]int
		source int
		target int
		count  int
	}{
		{routes: [][]int{{1, 2, 7}, {3, 6, 7}}, source: 1, target: 6, count: 2},
		{routes: [][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, source: 15, target: 12, count: -1},
	} {
		count := numBusesToDestination(tc.routes, tc.source, tc.target)
		assert.Equal(t, tc.count, count)
	}
}
