package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestInsertIntoBST$
func TestInsertIntoBST(t *testing.T) {
	for _, tc := range []struct {
		root    *TreeNode
		val     int
		newRoot *TreeNode
	}{
		{
			root:    makeTree(0, []any{4, 2, 7, 1, 3}),
			val:     5,
			newRoot: makeTree(0, []any{4, 2, 7, 1, 3, 5}),
		},
		{
			root:    makeTree(0, []any{40, 20, 60, 10, 30, 50, 70}),
			val:     25,
			newRoot: makeTree(0, []any{40, 20, 60, 10, 30, 50, 70, nil, nil, 25}),
		},
		{
			root:    makeTree(0, []any{4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil}),
			val:     5,
			newRoot: makeTree(0, []any{4, 2, 7, 1, 3, 5}),
		},
	} {
		newRoot := insertIntoBST(tc.root, tc.val)
		assert.Equal(t, tc.newRoot, newRoot)
	}
}