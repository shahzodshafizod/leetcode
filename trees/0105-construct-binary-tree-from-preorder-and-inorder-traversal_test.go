package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestBuildTree$
func TestBuildTree(t *testing.T) {
	for _, tc := range []struct {
		preorder []int
		inorder  []int
		root     *design.TreeNode
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			root:     design.MakeTree2(3, 9, 20, nil, nil, 15, 7),
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			root:     design.MakeTree2(-1),
		},
		{
			preorder: []int{3, 9, 10, 20, 15, 11, 7},
			inorder:  []int{10, 9, 3, 11, 15, 20, 7},
			root:     design.MakeTree2(3, 9, 20, 10, nil, 15, 7, nil, nil, 11),
		},
		{
			preorder: []int{3, 1, 2, 4},
			inorder:  []int{1, 2, 3, 4},
			root:     design.MakeTree2(3, 1, 4, nil, 2),
		},
	} {
		root := buildTree(tc.preorder, tc.inorder)
		assert.Equal(t, tc.root, root)
	}
}
