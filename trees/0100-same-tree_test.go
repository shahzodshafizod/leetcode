package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsSameTree$
func TestIsSameTree(t *testing.T) {
	for _, tc := range []struct {
		p    *TreeNode
		q    *TreeNode
		same bool
	}{
		{p: makeTree(0, []any{1, 2, 3}), q: makeTree(0, []any{1, 2, 3}), same: true},
		{p: makeTree(0, []any{1, 2}), q: makeTree(0, []any{1, nil, 2}), same: false},
		{p: makeTree(0, []any{1, 2, 1}), q: makeTree(0, []any{1, 1, 2}), same: false},
	} {
		same := isSameTree(tc.p, tc.q)
		assert.Equal(t, tc.same, same)
	}
}
