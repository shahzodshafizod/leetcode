package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestDelNodes$
func TestDelNodes(t *testing.T) {
	for _, tc := range []struct {
		root       *TreeNode
		toDelete   []int
		forestVals [][]any
		forest     []*TreeNode
	}{
		{
			root:       makeTree(0, []any{1, 2, 3, 4, 5, 6, 7}),
			toDelete:   []int{3, 5},
			forestVals: [][]any{{1, 2, nil, 4}, {6}, {7}},
		},
		{
			root:       makeTree(0, []any{1, 2, 4, nil, 3}),
			toDelete:   []int{3},
			forestVals: [][]any{{1, 2, 4}},
		},
	} {
		forest := delNodes(tc.root, tc.toDelete)
		tc.forest = make([]*TreeNode, len(tc.forestVals))
		for idx, vals := range tc.forestVals {
			tc.forest[idx] = makeTree(0, vals)
		}
		assert.Equal(t, tc.forest, forest)
	}
}
