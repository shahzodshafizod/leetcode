package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestMergeTrees$
func TestMergeTrees(t *testing.T) {
	for _, tc := range []struct {
		root1  *pkg.TreeNode
		root2  *pkg.TreeNode
		merged *pkg.TreeNode
	}{
		{root1: pkg.MakeTree2(1, 3, 2, 5), root2: pkg.MakeTree2(2, 1, 3, nil, 4, nil, 7), merged: pkg.MakeTree2(3, 4, 5, 5, 4, nil, 7)},
		{root1: pkg.MakeTree2(1), root2: pkg.MakeTree2(1, 2), merged: pkg.MakeTree2(2, 2)},
	} {
		merged := mergeTrees(tc.root1, tc.root2)
		assert.Equal(t, tc.merged, merged)
	}
}
