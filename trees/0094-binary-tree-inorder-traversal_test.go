package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestInorderTraversal$
func TestInorderTraversal(t *testing.T) {
	for _, tc := range []struct {
		root   *TreeNode
		values []int
	}{
		{root: makeTree(0, []any{1, nil, 2, nil, nil, 3}), values: []int{1, 3, 2}},
		{root: makeTree(0, []any{}), values: []int{}},
		{root: makeTree(0, []any{1}), values: []int{1}},
		{root: makeTree(0, []any{1, 2, 3, 4, 5, 6}), values: []int{4, 2, 5, 1, 6, 3}},
	} {
		values := inorderTraversal(tc.root)
		assert.Equal(t, tc.values, values)
	}
}
