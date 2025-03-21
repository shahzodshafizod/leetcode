package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestDiameterOfBinaryTree$
func TestDiameterOfBinaryTree(t *testing.T) {
	for _, tc := range []struct {
		root     *pkg.TreeNode
		diameter int
	}{
		{root: pkg.MakeTree(0, []any{1, 2, 3, 4, 5}), diameter: 3},
		{root: pkg.MakeTree(0, []any{1, 2}), diameter: 1},
		{root: pkg.MakeTree(0, []any{1, 2, nil, 3, 4, nil, nil, 5, nil, nil, 6}), diameter: 4},
	} {
		diameter := diameterOfBinaryTree(tc.root)
		assert.Equal(t, tc.diameter, diameter)
	}
}
