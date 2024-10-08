package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestLevelOrder$
func TestLevelOrder(t *testing.T) {
	for _, tc := range []struct {
		root   *pkg.TreeNode
		levels [][]int
	}{
		{root: pkg.MakeTree2(3, 9, 20, nil, nil, 15, 7), levels: [][]int{{3}, {9, 20}, {15, 7}}},
		{root: pkg.MakeTree2(1), levels: [][]int{{1}}},
		{root: pkg.MakeTree2(), levels: [][]int{}},
		{root: pkg.MakeTree2(3, 6, 1, 9, 2, nil, 4, nil, 5, nil, nil, nil, nil, 8), levels: [][]int{{3}, {6, 1}, {9, 2, 4}, {5}, {8}}},
		{root: pkg.MakeTree2(1, nil, 2, nil, 3, nil, 4, nil, 5), levels: [][]int{{1}, {2}, {3}, {4}, {5}}},
	} {
		levels := levelOrder(tc.root)
		assert.Equal(t, tc.levels, levels)
	}
}
