package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsValidBST$
func TestIsValidBST(t *testing.T) {
	for _, tc := range []struct {
		root  *TreeNode
		valid bool
	}{
		{root: makeTree(0, []any{12, 8, 18, 5, 10, 14, 25}), valid: true},
		{root: makeTree(0, []any{16, 8, 22, 9, nil, 19, 25}), valid: false},
		{root: makeTree(0, []any{13, 6, 17, 2, nil, 10, 22}), valid: false},
		{root: makeTree(0, []any{}), valid: true},
		{root: makeTree(0, []any{12, 7, 18, 5, 9, 16, 25}), valid: true},
		{root: makeTree(0, []any{1}), valid: true},
		{root: makeTree(0, []any{2, 1, 3}), valid: true},
		{root: makeTree(0, []any{5, 1, 4, nil, nil, 3, 6}), valid: false},
	} {
		valid := isValidBST(tc.root)
		assert.Equal(t, tc.valid, valid)
	}
}
