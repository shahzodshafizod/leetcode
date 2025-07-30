package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestCountPairs$
func TestCountPairs(t *testing.T) {
	for _, tc := range []struct {
		root      *pkg.TreeNode
		distance  int
		goodPairs int
	}{
		{root: pkg.MakeTree(0, []any{1, 2, 3, nil, 4}), distance: 3, goodPairs: 1},
		{root: pkg.MakeTree(0, []any{1, 2, 3, 4, 5, 6, 7}), distance: 3, goodPairs: 2},
		{root: pkg.MakeTree(0, []any{7, 1, 4, 6, nil, 5, 3, nil, nil, nil, nil, 2}), distance: 3, goodPairs: 1},
		{root: pkg.MakeTree(0, []any{1, 1, 1}), distance: 2, goodPairs: 1},
	} {
		goodPairs := countPairs(tc.root, tc.distance)
		assert.Equal(t, tc.goodPairs, goodPairs)
	}
}
