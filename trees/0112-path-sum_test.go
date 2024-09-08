package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestHasPathSum$
func TestHasPathSum(t *testing.T) {
	for _, tc := range []struct {
		root      *design.TreeNode
		targetSum int
		has       bool
	}{
		{
			root:      design.MakeTree(0, []any{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}),
			targetSum: 22,
			has:       true,
		},
		{
			root:      design.MakeTree(0, []any{1, 2, 3}),
			targetSum: 5,
			has:       false,
		},
		{
			root:      design.MakeTree(0, []any{}),
			targetSum: 0,
			has:       false,
		},
		{
			root:      design.MakeTree(0, []any{1, 2}),
			targetSum: 1,
			has:       false,
		},
		{
			root:      design.MakeTree(0, []any{1, 2, nil, 3, nil, 4, nil, 5}),
			targetSum: 6,
			has:       false,
		},
	} {
		has := hasPathSum(tc.root, tc.targetSum)
		assert.Equal(t, tc.has, has)
	}
}
