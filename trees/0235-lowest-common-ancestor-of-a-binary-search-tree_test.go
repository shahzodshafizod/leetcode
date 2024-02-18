package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestLowestCommonAncestor$
func TestLowestCommonAncestor(t *testing.T) {
	for _, tc := range []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		lca  *TreeNode
	}{
		{
			root: makeTree(0, []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 8},
			lca:  &TreeNode{Val: 6},
		},
		{
			root: makeTree(0, []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 4},
			lca:  &TreeNode{Val: 2},
		},
		{
			root: makeTree(0, []any{2, 1}),
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 1},
			lca:  &TreeNode{Val: 2},
		},
	} {
		lca := lowestCommonAncestor(tc.root, tc.p, tc.q)
		assert.Equal(t, tc.lca.Val, lca.Val)
	}
}
