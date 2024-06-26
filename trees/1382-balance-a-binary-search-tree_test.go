package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestBalanceBST$
func TestBalanceBST(t *testing.T) {
	for _, tc := range []struct {
		root     *TreeNode
		balanced *TreeNode
	}{
		{
			root:     makeTree(0, []any{1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4}),
			balanced: makeTree(0, []any{2, 1, 3, nil, nil, nil, 4}),
		},
		{root: makeTree(0, []any{2, 1, 3}), balanced: makeTree(0, []any{2, 1, 3})},
	} {
		balanced := balanceBST(tc.root)
		assert.Equal(t, tc.balanced, balanced)
	}
}
