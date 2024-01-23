package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestMaxDepth$
func TestMaxDepth(t *testing.T) {
	for _, tc := range []struct {
		root     *TreeNode
		maxDepth int
	}{
		{root: makeTree(0, []any{}), maxDepth: 0},
		{root: makeTree(0, []any{1}), maxDepth: 1},
		{root: makeTree(0, []any{3, 9, 20, nil, nil, 15, 7}), maxDepth: 3},
		{root: makeTree(0, []any{1, nil, 2}), maxDepth: 2},
		{root: makeTree(0, []any{1, 2, 3, 4, 5, nil, nil, nil, 6, nil, nil, nil, nil, nil, nil, nil, nil, nil, 7, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}), maxDepth: 5},
		{root: makeTree(0, []any{1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 5}), maxDepth: 5},
	} {
		maxDepth := maxDepth(tc.root)
		assert.Equal(t, tc.maxDepth, maxDepth)
	}
}
