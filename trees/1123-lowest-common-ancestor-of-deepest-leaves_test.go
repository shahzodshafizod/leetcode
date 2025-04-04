package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestLcaDeepestLeaves$
func TestLcaDeepestLeaves(t *testing.T) {
	for _, tc := range []struct {
		root     *pkg.TreeNode
		ancestor *pkg.TreeNode
	}{
		{root: pkg.MakeTree2(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4), ancestor: pkg.MakeTree2(2, 7, 4)},
		{root: pkg.MakeTree2(1), ancestor: pkg.MakeTree2(1)},
		{root: pkg.MakeTree2(0, 1, 3, nil, 2), ancestor: pkg.MakeTree2(2)},
	} {
		ancestor := lcaDeepestLeaves(tc.root)
		assert.Equal(t, tc.ancestor, ancestor)
	}
}
