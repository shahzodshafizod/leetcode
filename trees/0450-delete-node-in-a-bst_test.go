package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestDeleteNode$
func TestDeleteNode(t *testing.T) {
	for _, tc := range []struct {
		root    *pkg.TreeNode
		key     int
		newRoot *pkg.TreeNode
	}{
		{
			root:    pkg.MakeTree(0, []any{5, 3, 6, 2, 4, nil, 7}),
			key:     3,
			newRoot: pkg.MakeTree(0, []any{5, 4, 6, 2, nil, nil, 7}),
		},
		{
			root:    pkg.MakeTree(0, []any{5, 3, 6, 2, 4, nil, 7}),
			key:     0,
			newRoot: pkg.MakeTree(0, []any{5, 3, 6, 2, 4, nil, 7}),
		},
		{
			root:    pkg.MakeTree(0, []any{}),
			key:     0,
			newRoot: pkg.MakeTree(0, []any{}),
		},
		{
			root:    pkg.MakeTree(0, []any{50, 30, 60, 20, 40, nil, 70, nil, nil, 35, 45}),
			key:     30,
			newRoot: pkg.MakeTree(0, []any{50, 35, 60, 20, 40, nil, 70, nil, nil, nil, 45}),
		},
	} {
		newRoot := deleteNode(tc.root, tc.key)
		assert.Equal(t, tc.newRoot, newRoot)
	}
}
