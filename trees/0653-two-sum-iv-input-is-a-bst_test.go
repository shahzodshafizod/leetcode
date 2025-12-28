package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestFindTarget$
func TestFindTarget(t *testing.T) {
	for _, tc := range []struct {
		root   *pkg.TreeNode
		k      int
		target bool
	}{
		{root: pkg.MakeTree(0, []any{5, 3, 6, 2, 4, nil, 7}), k: 9, target: true},   // 3 + 6 = 9
		{root: pkg.MakeTree(0, []any{5, 3, 6, 2, 4, nil, 7}), k: 28, target: false}, // No pair sums to 28
		{root: pkg.MakeTree(0, []any{2, 1, 3}), k: 4, target: true},                 // 1 + 3 = 4
		{root: pkg.MakeTree(0, []any{2, 1, 3}), k: 1, target: false},                // No pair sums to 1
		{root: pkg.MakeTree(0, []any{2, 1, 3}), k: 3, target: true},                 // 1 + 2 = 3
		{root: pkg.MakeTree(0, []any{1}), k: 2, target: false},                      // Single node, no pair
	} {
		target := findTarget(tc.root, tc.k)
		assert.Equal(t, tc.target, target)
	}
}
