package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestRemoveLeafNodes$
func TestRemoveLeafNodes(t *testing.T) {
	for _, tc := range []struct {
		root    *pkg.TreeNode
		target  int
		newroot *pkg.TreeNode
	}{
		{root: pkg.MakeTree(0, []any{1, 2, 3, 2, nil, 2, 4}), target: 2, newroot: pkg.MakeTree(0, []any{1, nil, 3, nil, nil, nil, 4})},
		{root: pkg.MakeTree(0, []any{1, 3, 3, 3, 2}), target: 3, newroot: pkg.MakeTree(0, []any{1, 3, nil, nil, 2})},
		{root: pkg.MakeTree(0, []any{1, 2, nil, 2, nil, 2}), target: 2, newroot: pkg.MakeTree(0, []any{1})},
	} {
		newroot := removeLeafNodes(tc.root, tc.target)
		assert.Equal(t, tc.newroot, newroot)
	}
}
