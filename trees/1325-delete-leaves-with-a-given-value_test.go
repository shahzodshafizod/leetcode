package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestRemoveLeafNodes$
func TestRemoveLeafNodes(t *testing.T) {
	for _, tc := range []struct {
		root    *design.TreeNode
		target  int
		newroot *design.TreeNode
	}{
		{root: design.MakeTree(0, []any{1, 2, 3, 2, nil, 2, 4}), target: 2, newroot: design.MakeTree(0, []any{1, nil, 3, nil, nil, nil, 4})},
		{root: design.MakeTree(0, []any{1, 3, 3, 3, 2}), target: 3, newroot: design.MakeTree(0, []any{1, 3, nil, nil, 2})},
		{root: design.MakeTree(0, []any{1, 2, nil, 2, nil, 2}), target: 2, newroot: design.MakeTree(0, []any{1})},
	} {
		newroot := removeLeafNodes(tc.root, tc.target)
		assert.Equal(t, tc.newroot, newroot)
	}
}
