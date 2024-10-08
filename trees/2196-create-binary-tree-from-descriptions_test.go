package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestCreateBinaryTree$
func TestCreateBinaryTree(t *testing.T) {
	for _, tc := range []struct {
		descriptions [][]int
		root         *pkg.TreeNode
	}{
		{
			descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			root:         pkg.MakeTree(0, []any{50, 20, 80, 15, 17, 19}),
		},
		{
			descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
			root:         pkg.MakeTree(0, []any{1, 2, nil, nil, 3, nil, nil, nil, nil, 4}),
		},
	} {
		root := createBinaryTree(tc.descriptions)
		assert.Equal(t, tc.root, root)
	}
}
