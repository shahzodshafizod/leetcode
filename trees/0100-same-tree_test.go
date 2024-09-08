package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsSameTree$
func TestIsSameTree(t *testing.T) {
	for _, tc := range []struct {
		p    *design.TreeNode
		q    *design.TreeNode
		same bool
	}{
		{p: design.MakeTree(0, []any{1, 2, 3}), q: design.MakeTree(0, []any{1, 2, 3}), same: true},
		{p: design.MakeTree(0, []any{1, 2}), q: design.MakeTree(0, []any{1, nil, 2}), same: false},
		{p: design.MakeTree(0, []any{1, 2, 1}), q: design.MakeTree(0, []any{1, 1, 2}), same: false},
	} {
		same := isSameTree(tc.p, tc.q)
		assert.Equal(t, tc.same, same)
	}
}
