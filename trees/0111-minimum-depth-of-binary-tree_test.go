package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestMinDepth$
func TestMinDepth(t *testing.T) {
	for _, tc := range []struct {
		root  *pkg.TreeNode
		depth int
	}{
		{root: pkg.MakeTree2(), depth: 0},
		{root: pkg.MakeTree2(2), depth: 1},
		{root: pkg.MakeTree2(3, 9, 20, nil, nil, 15, 7), depth: 2},
		{root: pkg.MakeTree2(2, nil, 3, nil, 4, nil, 5, nil, 6), depth: 5},
	} {
		depth := minDepth(tc.root)
		assert.Equal(t, tc.depth, depth)
	}
}
