package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestSortedArrayToBST$
func TestSortedArrayToBST(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		root *pkg.TreeNode
	}{
		{nums: []int{-10, -3, 0, 5, 9}, root: pkg.MakeTree2(0, -3, 9, -10, nil, 5)},
		{nums: []int{1, 3}, root: pkg.MakeTree2(3, 1)},
	} {
		root := sortedArrayToBST(tc.nums)
		assert.Equal(t, tc.root, root)
	}
}
