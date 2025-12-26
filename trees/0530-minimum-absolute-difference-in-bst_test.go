package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestGetMinimumDifference$
func TestGetMinimumDifference(t *testing.T) {
	for _, tc := range []struct {
		root    *pkg.TreeNode
		minDiff int
	}{
		{root: pkg.MakeTree2(4, 2, 6, 1, 3), minDiff: 1},
		{root: pkg.MakeTree2(1, 0, 48, nil, nil, 12, 49), minDiff: 1},
		{root: pkg.MakeTree2(5, 4, 7), minDiff: 1},
	} {
		minDiff := getMinimumDifference(tc.root)
		assert.Equal(t, tc.minDiff, minDiff)
	}
}
