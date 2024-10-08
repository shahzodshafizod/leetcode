package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestMaxDepth$
func TestMaxDepth(t *testing.T) {
	for _, tc := range []struct {
		root     *pkg.TreeNode
		maxDepth int
	}{
		{root: pkg.MakeTree2(3, 9, 20, nil, nil, 15, 7), maxDepth: 3},
		{root: pkg.MakeTree2(1, nil, 2), maxDepth: 2},
		{root: pkg.MakeTree2(), maxDepth: 0},
		{root: pkg.MakeTree2(1), maxDepth: 1},
		{root: pkg.MakeTree2(1, 2, 3, 4, 5, nil, nil, nil, 3, nil, nil, nil, 7), maxDepth: 5},
		{root: pkg.MakeTree2(1, nil, 2, nil, 3, nil, 4, nil, 5), maxDepth: 5},
	} {
		maxDepth := maxDepth(tc.root)
		assert.Equal(t, tc.maxDepth, maxDepth)
	}
}
