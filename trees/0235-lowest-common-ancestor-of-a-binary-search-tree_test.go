package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestLowestCommonAncestor$
func TestLowestCommonAncestor(t *testing.T) {
	for _, tc := range []struct {
		root *pkg.TreeNode
		p    *pkg.TreeNode
		q    *pkg.TreeNode
		lca  *pkg.TreeNode
	}{
		{
			root: pkg.MakeTree(0, []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			p:    &pkg.TreeNode{Val: 2},
			q:    &pkg.TreeNode{Val: 8},
			lca:  &pkg.TreeNode{Val: 6},
		},
		{
			root: pkg.MakeTree(0, []any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			p:    &pkg.TreeNode{Val: 2},
			q:    &pkg.TreeNode{Val: 4},
			lca:  &pkg.TreeNode{Val: 2},
		},
		{
			root: pkg.MakeTree(0, []any{2, 1}),
			p:    &pkg.TreeNode{Val: 2},
			q:    &pkg.TreeNode{Val: 1},
			lca:  &pkg.TreeNode{Val: 2},
		},
	} {
		lca := lowestCommonAncestor(tc.root, tc.p, tc.q)
		assert.Equal(t, tc.lca.Val, lca.Val)
	}
}
