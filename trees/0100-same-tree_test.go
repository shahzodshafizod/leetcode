package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsSameTree$
func TestIsSameTree(t *testing.T) {
	for _, tc := range []struct {
		p    *pkg.TreeNode
		q    *pkg.TreeNode
		same bool
	}{
		{p: pkg.MakeTree2(1, 2, 3), q: pkg.MakeTree2(1, 2, 3), same: true},
		{p: pkg.MakeTree2(1, 2), q: pkg.MakeTree2(1, nil, 2), same: false},
		{p: pkg.MakeTree2(1, 2, 1), q: pkg.MakeTree2(1, 1, 2), same: false},
	} {
		same := isSameTree(tc.p, tc.q)
		assert.Equal(t, tc.same, same)
	}
}
