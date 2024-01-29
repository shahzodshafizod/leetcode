package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestTraversalBFS$
func TestTraversalBFS(t *testing.T) {
	for _, tc := range []struct {
		root   *TreeNode
		values []int
	}{
		{root: makeTree(0, []any{}), values: []int{}},
	} {
		values := traversalBFS(tc.root)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./trees/ -run ^TestTraversalDFS$
func TestTraversalDFS(t *testing.T) {
	for _, tc := range []struct {
		root   *TreeNode
		values []int
	}{
		{root: makeTree(0, []any{}), values: []int{}},
	} {
		values := traversalDFS(tc.root)
		assert.Equal(t, tc.values, values)
	}
}
