package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestAddOneRow$
func TestAddOneRow(t *testing.T) {
	for _, tc := range []struct {
		root    *pkg.TreeNode
		val     int
		depth   int
		newroot *pkg.TreeNode
	}{
		{root: pkg.MakeTree(0, []any{4, 2, 6, 3, 1, 5}), val: 1, depth: 2, newroot: pkg.MakeTree(0, []any{4, 1, 1, 2, nil, nil, 6, 3, 1, nil, nil, nil, nil, 5})},
		{root: pkg.MakeTree(0, []any{4, 2, nil, 3, 1}), val: 1, depth: 3, newroot: pkg.MakeTree(0, []any{4, 2, nil, 1, 1, nil, nil, 3, nil, nil, 1})},
		{root: pkg.MakeTree(0, []any{1}), val: 2, depth: 1, newroot: pkg.MakeTree(0, []any{2, 1})},
		{root: pkg.MakeTree(0, []any{1, 2, 3, 4}), val: 5, depth: 4, newroot: pkg.MakeTree(0, []any{1, 2, 3, 4, nil, nil, nil, 5, 5})},
	} {
		newroot := addOneRow(tc.root, tc.val, tc.depth)
		assert.Equal(t, tc.newroot, newroot)
	}
}
