package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestSmallestFromLeaf$
func TestSmallestFromLeaf(t *testing.T) {
	for _, tc := range []struct {
		root     *pkg.TreeNode
		smallest string
	}{
		{root: pkg.MakeTree(0, []any{0, 1, 2, 3, 4, 3, 4}), smallest: "dba"},
		{root: pkg.MakeTree(0, []any{25, 1, 3, 1, 3, 0, 2}), smallest: "adz"},
		{root: pkg.MakeTree(0, []any{2, 2, 1, nil, 1, 0, nil, 0}), smallest: "abc"},
		{root: pkg.MakeTree(0, []any{25, 1, nil, 0, 0, nil, nil, 1, nil, nil, nil, nil, nil, nil, nil, 0}), smallest: "ababz"},
		{root: pkg.MakeTree(0, []any{4, 0, 1, 1}), smallest: "bae"},
		{root: pkg.MakeTree(0, []any{3, 9, 20, nil, nil, 15, 7}), smallest: "hud"},
		{root: pkg.MakeTree(0, []any{2, 1, 1, nil, 1, 0, nil, nil, nil, 0}), smallest: "abbc"},
	} {
		smallest := smallestFromLeaf(tc.root)
		assert.Equal(t, tc.smallest, smallest)
	}
}
