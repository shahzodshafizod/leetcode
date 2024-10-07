package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestLowestCommonAncestor$
func TestLowestCommonAncestor(t *testing.T) {
	for _, tc := range []struct {
		root *design.TreeNode
		p    *design.TreeNode
		q    *design.TreeNode
		lca  *design.TreeNode
	}{
		{
			root: design.MakeTree(0, []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			p:    &design.TreeNode{Val: 2},
			q:    &design.TreeNode{Val: 8},
			lca:  &design.TreeNode{Val: 6},
		},
		{
			root: design.MakeTree(0, []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			p:    &design.TreeNode{Val: 2},
			q:    &design.TreeNode{Val: 4},
			lca:  &design.TreeNode{Val: 2},
		},
		{
			root: design.MakeTree(0, []any{2, 1}),
			p:    &design.TreeNode{Val: 2},
			q:    &design.TreeNode{Val: 1},
			lca:  &design.TreeNode{Val: 2},
		},
	} {
		lca := lowestCommonAncestor(tc.root, tc.p, tc.q)
		assert.Equal(t, tc.lca.Val, lca.Val)
	}
}
