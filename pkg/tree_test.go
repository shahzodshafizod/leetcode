package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./pkg/ -run ^TestTraversalBFS$
func TestTraversalBFS(t *testing.T) {
	for _, tc := range []struct {
		root   *TreeNode
		values []int
	}{
		{root: MakeTree(0, []any{9, 4, 20, 1, 6, 15, 170}), values: []int{9, 4, 20, 1, 6, 15, 170}},
	} {
		values := traversalBFS(tc.root)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./pkg/ -run ^TestTraversalDFS$
func TestTraversalDFS(t *testing.T) {
	for _, tc := range []struct {
		root   *TreeNode
		values []int
	}{
		{
			root:   MakeTree(0, []any{9, 4, 20, 1, 6, 15, 170}),
			values: []int{1, 4, 6, 9, 15, 20, 170},
		},
	} {
		values := traversalDFSInOrder(tc.root)
		assert.Equal(t, tc.values, values)
	}
}

// go test -v -count=1 ./pkg/ -run ^TestTraversalMorris$
func TestTraversalMorris(t *testing.T) {
	for _, tc := range []struct {
		root   *TreeNode
		values []int
	}{
		{
			root:   MakeTree(0, []any{9, 4, 20, 1, 6, 15, 170}),
			values: []int{1, 4, 6, 9, 15, 20, 170},
		},
	} {
		values := traversalMorris(tc.root)
		assert.Equal(t, tc.values, values)
	}
}
