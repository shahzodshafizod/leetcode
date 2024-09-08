package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestBalanceBST$
func TestBalanceBST(t *testing.T) {
	for _, tc := range []struct {
		root     *design.TreeNode
		balanced *design.TreeNode
	}{
		{
			root:     design.MakeTree(0, []any{1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4}),
			balanced: design.MakeTree(0, []any{2, 1, 3, nil, nil, nil, 4}),
		},
		{root: design.MakeTree(0, []any{2, 1, 3}), balanced: design.MakeTree(0, []any{2, 1, 3})},
	} {
		balanced := balanceBST(tc.root)
		assert.Equal(t, tc.balanced, balanced)
	}
}
