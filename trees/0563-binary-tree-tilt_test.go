package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestFindTilt$
func TestFindTilt(t *testing.T) {
	for _, tc := range []struct {
		root *pkg.TreeNode
		tilt int
	}{
		{root: pkg.MakeTree(0, []any{1, 2, 3}), tilt: 1},
		{root: pkg.MakeTree(0, []any{4, 2, 9, 3, 5, nil, 7}), tilt: 15},
		{root: pkg.MakeTree(0, []any{21, 7, 14, 1, 1, 2, 2, 3, 3}), tilt: 9},
		{root: pkg.MakeTree(0, []any{}), tilt: 0},
		{root: pkg.MakeTree(0, []any{1}), tilt: 0},
	} {
		tilt := findTilt(tc.root)
		assert.Equal(t, tc.tilt, tilt)
	}
}
