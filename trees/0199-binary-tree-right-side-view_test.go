package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestRightSideView$
func TestRightSideView(t *testing.T) {
	for _, tc := range []struct {
		root   *TreeNode
		values []int
	}{
		{root: makeTree(0, []any{1, 2, 3, 4, 5, nil, 6, nil, 7, nil, nil, nil, nil, nil, nil, nil, nil, 8, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}), values: []int{1, 3, 6, 7, 8}},
		{root: makeTree(0, []any{1}), values: []int{1}},
		{root: makeTree(0, []any{}), values: []int{}},
		{root: makeTree(0, []any{1, 2, 3, nil, 5, nil, 4}), values: []int{1, 3, 4}},
		{root: makeTree(0, []any{1, nil, 3}), values: []int{1, 3}},
	} {
		values := rightSideView(tc.root)
		assert.Equal(t, tc.values, values)
	}
}
