package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsSubtree$
func TestIsSubtree(t *testing.T) {
	for _, tc := range []struct {
		root    *pkg.TreeNode
		subRoot *pkg.TreeNode
		is      bool
	}{
		{root: pkg.MakeTree(0, []any{3, 4, 5, 1, 2}), subRoot: pkg.MakeTree(0, []any{4, 1, 2}), is: true},
		{root: pkg.MakeTree(0, []any{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}), subRoot: pkg.MakeTree(0, []any{4, 1, 2}), is: false},
		{root: pkg.MakeTree(0, []any{1, 1}), subRoot: pkg.MakeTree(0, []any{1}), is: true},
		{root: pkg.MakeTree(0, []any{1, 2, 3}), subRoot: pkg.MakeTree(0, []any{1, 2}), is: false},
		{root: pkg.MakeTree(0, []any{4, nil, 2}), subRoot: pkg.MakeTree(0, []any{3, 4, 5, nil, 2, nil, 2}), is: false},
		{root: pkg.MakeTree(0, []any{3, 4, 5, nil, 2, nil, 2}), subRoot: pkg.MakeTree(0, []any{4, nil, 2}), is: true},
		{root: pkg.MakeTree(0, []any{3, 1, 2, 1, nil, 2}), subRoot: pkg.MakeTree(0, []any{3, 1, 2}), is: false},
		{root: pkg.MakeTree(0, []any{1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, 2}), subRoot: pkg.MakeTree(0, []any{1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, 2}), is: true},
		{root: pkg.MakeTree(0, []any{3, 4, 5, 1, nil, 2}), subRoot: pkg.MakeTree(0, []any{3, 1, 2}), is: false},
		{root: pkg.MakeTree(0, []any{4, 1}), subRoot: pkg.MakeTree(0, []any{4, 1, 2}), is: false},
		{root: pkg.MakeTree(0, []any{4, 1, 2}), subRoot: pkg.MakeTree(0, []any{4, 1}), is: false},
	} {
		is := isSubtree(tc.root, tc.subRoot)
		assert.Equal(t, tc.is, is)
	}
}
