package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestBstToGst$
func TestBstToGst(t *testing.T) {
	for _, tc := range []struct {
		rootBST *pkg.TreeNode
		rootGST *pkg.TreeNode
	}{
		{
			rootBST: pkg.MakeTree(0, []any{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}),
			rootGST: pkg.MakeTree(0, []any{30, 36, 21, 36, 35, 26, 15, nil, nil, nil, 33, nil, nil, nil, 8}),
		},
		{
			rootBST: pkg.MakeTree(0, []any{0, nil, 1}),
			rootGST: pkg.MakeTree(0, []any{1, nil, 1}),
		},
	} {
		rootGST := bstToGst(tc.rootBST)
		assert.Equal(t, tc.rootGST, rootGST)
	}
}
