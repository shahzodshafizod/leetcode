package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestInsertIntoBST$
func TestInsertIntoBST(t *testing.T) {
	for _, tc := range []struct {
		root    *design.TreeNode
		val     int
		newRoot *design.TreeNode
	}{
		{
			root:    design.MakeTree(0, []any{4, 2, 7, 1, 3}),
			val:     5,
			newRoot: design.MakeTree(0, []any{4, 2, 7, 1, 3, 5}),
		},
		{
			root:    design.MakeTree(0, []any{40, 20, 60, 10, 30, 50, 70}),
			val:     25,
			newRoot: design.MakeTree(0, []any{40, 20, 60, 10, 30, 50, 70, nil, nil, 25}),
		},
		{
			root:    design.MakeTree(0, []any{4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil}),
			val:     5,
			newRoot: design.MakeTree(0, []any{4, 2, 7, 1, 3, 5}),
		},
	} {
		newRoot := insertIntoBST(tc.root, tc.val)
		assert.Equal(t, tc.newRoot, newRoot)
	}
}
