package graphs

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCloneGraph$
func TestCloneGraph(t *testing.T) {
	for _, tc := range []struct {
		node  *pkg.Node
		clone *pkg.Node
	}{
		{node: pkg.MakeGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}), clone: pkg.MakeGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})},
		{node: pkg.MakeGraph([][]int{{}}), clone: pkg.MakeGraph([][]int{{}})},
		{node: pkg.MakeGraph([][]int{}), clone: pkg.MakeGraph([][]int{})},
	} {
		clone := cloneGraph(tc.node)
		assert.Equal(t, tc.clone, clone)
	}
}
