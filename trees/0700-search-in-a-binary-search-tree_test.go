package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestSearchBST$
func TestSearchBST(t *testing.T) {
	for _, tc := range []struct {
		root *pkg.TreeNode
		val  int
		node *pkg.TreeNode
	}{
		{root: pkg.MakeTree(0, []any{4, 2, 7, 1, 3}), val: 2, node: pkg.MakeTree(0, []any{2, 1, 3})},
		{root: pkg.MakeTree(0, []any{4, 2, 7, 1, 3}), val: 5, node: pkg.MakeTree(0, []any{})},
	} {
		node := searchBST(tc.root, tc.val)
		assert.Equal(t, tc.node, node)
	}
}
