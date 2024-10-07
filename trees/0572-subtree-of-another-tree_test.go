package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsSubtree$
func TestIsSubtree(t *testing.T) {
	for _, tc := range []struct {
		root    *design.TreeNode
		subRoot *design.TreeNode
		is      bool
	}{
		{root: design.MakeTree(0, []any{3, 4, 5, 1, 2}), subRoot: design.MakeTree(0, []any{4, 1, 2}), is: true},
		{root: design.MakeTree(0, []any{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}), subRoot: design.MakeTree(0, []any{4, 1, 2}), is: false},
		{root: design.MakeTree(0, []any{1, 1}), subRoot: design.MakeTree(0, []any{1}), is: true},
		{root: design.MakeTree(0, []any{1, 2, 3}), subRoot: design.MakeTree(0, []any{1, 2}), is: false},
		{root: design.MakeTree(0, []any{4, nil, 2}), subRoot: design.MakeTree(0, []any{3, 4, 5, nil, 2, nil, 2}), is: false},
		{root: design.MakeTree(0, []any{3, 4, 5, nil, 2, nil, 2}), subRoot: design.MakeTree(0, []any{4, nil, 2}), is: true},
		{root: design.MakeTree(0, []any{3, 1, 2, 1, nil, 2}), subRoot: design.MakeTree(0, []any{3, 1, 2}), is: false},
		{root: design.MakeTree(0, []any{1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, 2}), subRoot: design.MakeTree(0, []any{1, nil, 1, nil, 1, nil, 1, nil, 1, nil, 1, 2}), is: true},
		{root: design.MakeTree(0, []any{3, 4, 5, 1, nil, 2}), subRoot: design.MakeTree(0, []any{3, 1, 2}), is: false},
		{root: design.MakeTree(0, []any{4, 1}), subRoot: design.MakeTree(0, []any{4, 1, 2}), is: false},
		{root: design.MakeTree(0, []any{4, 1, 2}), subRoot: design.MakeTree(0, []any{4, 1}), is: false},
	} {
		is := isSubtree(tc.root, tc.subRoot)
		assert.Equal(t, tc.is, is)
	}
}
