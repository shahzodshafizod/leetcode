package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestKthSmallest$
func TestKthSmallest(t *testing.T) {
	for _, tc := range []struct {
		root *pkg.TreeNode
		k    int
		kth  int
	}{
		{root: pkg.MakeTree(0, []any{3, 1, 4, nil, 2}), k: 1, kth: 1},
		{root: pkg.MakeTree(0, []any{5, 3, 6, 2, 4, nil, nil, 1}), k: 3, kth: 3},
	} {
		kth := kthSmallest(tc.root, tc.k)
		assert.Equal(t, tc.kth, kth)
	}
}
