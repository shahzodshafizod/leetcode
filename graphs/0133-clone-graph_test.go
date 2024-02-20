package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCloneGraph$
func TestCloneGraph(t *testing.T) {
	for _, tc := range []struct {
		node  *Node
		clone *Node
	}{
		{node: makeGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}), clone: makeGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})},
		{node: makeGraph([][]int{{}}), clone: makeGraph([][]int{{}})},
		{node: makeGraph([][]int{}), clone: makeGraph([][]int{})},
	} {
		clone := cloneGraph(tc.node)
		assert.Equal(t, tc.clone, clone)
	}
}
