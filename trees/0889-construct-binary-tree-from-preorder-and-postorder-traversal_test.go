package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestConstructFromPrePost$
func TestConstructFromPrePost(t *testing.T) {
	for _, tc := range []struct {
		preorder  []int
		postorder []int
		root      *pkg.TreeNode
	}{
		{preorder: []int{1, 2, 4, 5, 3, 6, 7}, postorder: []int{4, 5, 2, 6, 7, 3, 1}, root: pkg.MakeTree2(1, 2, 3, 4, 5, 6, 7)},
		{preorder: []int{1}, postorder: []int{1}, root: pkg.MakeTree2(1)},
	} {
		root := constructFromPrePost(tc.preorder, tc.postorder)
		assert.Equal(t, tc.root, root)
	}
}
