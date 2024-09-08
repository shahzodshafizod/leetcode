package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestSumOfLeftLeaves$
func TestSumOfLeftLeaves(t *testing.T) {
	for _, tc := range []struct {
		root *design.TreeNode
		sum  int
	}{
		{root: design.MakeTree(0, []any{3, 9, 20, nil, nil, 15, 7}), sum: 24},
		{root: design.MakeTree(0, []any{1}), sum: 0},
		{root: design.MakeTree(0, []any{1, 2, 3, 4, 5}), sum: 4},
		{root: design.MakeTree(0, []any{1, 3, 2, nil, 2}), sum: 0},
		{root: design.MakeTree(0, []any{3, 9, 20, nil, nil, 3, 9, 20, nil, nil, 15, 7, 15, 7}), sum: 23},
	} {
		sum := sumOfLeftLeaves(tc.root)
		assert.Equal(t, tc.sum, sum)
	}
}
