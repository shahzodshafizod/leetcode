package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestInsertIntoBST$
func TestInsertIntoBST(t *testing.T) {
	for _, tc := range []struct {
		root    *pkg.TreeNode
		val     int
		newRoot *pkg.TreeNode
	}{
		{
			root:    pkg.MakeTree(0, []any{4, 2, 7, 1, 3}),
			val:     5,
			newRoot: pkg.MakeTree(0, []any{4, 2, 7, 1, 3, 5}),
		},
		{
			root:    pkg.MakeTree(0, []any{40, 20, 60, 10, 30, 50, 70}),
			val:     25,
			newRoot: pkg.MakeTree(0, []any{40, 20, 60, 10, 30, 50, 70, nil, nil, 25}),
		},
		{
			root:    pkg.MakeTree(0, []any{4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil}),
			val:     5,
			newRoot: pkg.MakeTree(0, []any{4, 2, 7, 1, 3, 5}),
		},
	} {
		newRoot := insertIntoBST(tc.root, tc.val)
		assert.Equal(t, tc.newRoot, newRoot)
	}
}
