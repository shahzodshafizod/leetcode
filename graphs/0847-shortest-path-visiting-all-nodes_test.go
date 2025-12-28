package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestShortestPathLength$
func TestShortestPathLength(t *testing.T) {
	for _, tc := range []struct {
		graph  [][]int
		length int
	}{
		{
			graph:  [][]int{{1, 2, 3}, {0}, {0}, {0}},
			length: 4,
		},
		{
			graph:  [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}},
			length: 4,
		},
		{
			graph:  [][]int{{}},
			length: 0,
		},
		{
			graph:  [][]int{{1}, {0}},
			length: 1,
		},
		{
			graph:  [][]int{{1, 2}, {0, 2}, {0, 1}},
			length: 2,
		},
	} {
		length := shortestPathLength(tc.graph)
		assert.Equal(t, tc.length, length)
	}
}
