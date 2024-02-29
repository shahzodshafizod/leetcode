package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsEvenOddTree$
func TestIsEvenOddTree(t *testing.T) {
	for _, tc := range []struct {
		root *TreeNode
		is   bool
	}{
		{root: makeTree(0, []any{1, 10, 4, 3, nil, 7, 9, 12, 8, 6, nil, nil, 2}), is: true},
		{root: makeTree(0, []any{5, 4, 2, 3, 3, 7}), is: false},
		{root: makeTree(0, []any{5, 9, 1, 3, 5, 7}), is: false},
	} {
		is := isEvenOddTree(tc.root)
		assert.Equal(t, tc.is, is)
	}
}
