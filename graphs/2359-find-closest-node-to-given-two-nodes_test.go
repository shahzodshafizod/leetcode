package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestClosestMeetingNode$
func TestClosestMeetingNode(t *testing.T) {
	for _, tc := range []struct {
		edges []int
		node1 int
		node2 int
		node  int
	}{
		{edges: []int{2, 2, 3, -1}, node1: 0, node2: 1, node: 2},
		{edges: []int{1, 2, -1}, node1: 0, node2: 2, node: 2},
		{edges: []int{-1, -1}, node1: 0, node2: 0, node: 0},
		{edges: []int{2, 0, 0}, node1: 2, node2: 0, node: 0},
	} {
		node := closestMeetingNode(tc.edges, tc.node1, tc.node2)
		assert.Equal(t, tc.node, node)
	}
}
