package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestBinaryTreePaths$
func TestBinaryTreePaths(t *testing.T) {
	for _, tc := range []struct {
		root  *pkg.TreeNode
		paths []string
	}{
		{root: pkg.MakeTree2(1, 2, 3, nil, 5), paths: []string{"1->2->5", "1->3"}},
		{root: pkg.MakeTree2(1), paths: []string{"1"}},
	} {
		paths := binaryTreePaths(tc.root)
		assert.Equal(t, tc.paths, paths)
	}
}
