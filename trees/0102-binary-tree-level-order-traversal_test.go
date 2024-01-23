package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestLevelOrder$
func TestLevelOrder(t *testing.T) {

	for _, tc := range []struct {
		root   *TreeNode
		levels [][]int
	}{
		{root: makeTree(0, []any{3, 9, 20, nil, nil, 15, 7}), levels: [][]int{{3}, {9, 20}, {15, 7}}},
		{root: makeTree(0, []any{1}), levels: [][]int{{1}}},
		{root: makeTree(0, []any{}), levels: [][]int{}},
		{root: makeTree(0, []any{3, 6, 1, 9, 2, nil, 4, nil, 5, nil, nil, nil, nil, nil, nil, nil, nil, 8, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}), levels: [][]int{{3}, {6, 1}, {9, 2, 4}, {5}, {8}}},
		{root: makeTree(0, []any{1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 5}), levels: [][]int{{1}, {2}, {3}, {4}, {5}}},
	} {
		levels := levelOrder(tc.root)
		assert.Equal(t, tc.levels, levels)
	}
}
