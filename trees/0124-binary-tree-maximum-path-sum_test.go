package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestMaxPathSum$
func TestMaxPathSum(t *testing.T) {
	for _, tc := range []struct {
		root    *design.TreeNode
		pathsum int
	}{
		{root: design.MakeTree2(1, 2, 3), pathsum: 6},
		{root: design.MakeTree2(-10, 9, 20, nil, nil, 15, 7), pathsum: 42},
		{root: design.MakeTree2(-10, -9, 20, nil, nil, -15, -7), pathsum: 20},
		{root: design.MakeTree2(2, -1), pathsum: 2},
		{root: design.MakeTree2(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1), pathsum: 48},
		{root: design.MakeTree2(9, 6, -3, nil, nil, -6, 2, nil, nil, 2, nil, -6, -6, -6), pathsum: 16},
		{root: design.MakeTree2(1, -2, -3, 1, 3, -2, nil, -1), pathsum: 3},
		{root: design.MakeTree2(0), pathsum: 0},
		{root: design.MakeTree2(1, 1, 1, 1, 1, 1, 1), pathsum: 5},
		{root: design.MakeTree2(1, 2, 3, 4, 5), pathsum: 11},
		{root: design.MakeTree2(1, 2, 1, 10, 10, 1, 1), pathsum: 22},
		{root: design.MakeTree2(-3), pathsum: -3},
		{root: design.MakeTree2(-1, nil, 9, -6, 3, nil, nil, nil, -2), pathsum: 12},
		{root: design.MakeTree2(-1, nil, 9, -6, 3, nil, nil, nil, 2), pathsum: 14},
		{root: design.MakeTree2(-1, 8, 2), pathsum: 9},
		{root: design.MakeTree2(-1, -2, 10, -6, nil, -3, -6), pathsum: 10},
		{root: design.MakeTree2(8, 9, -6, nil, nil, 5, 9), pathsum: 20},
		{root: design.MakeTree2(-1, 8, 2, nil, -9, 0, nil, nil, nil, -3, nil, nil, -9, nil, 2), pathsum: 9},
		{root: design.MakeTree2(2, -1, -2), pathsum: 2},
		{root: design.MakeTree2(10, -20, -30), pathsum: 10},
		{root: design.MakeTree2(1, -2, 3), pathsum: 4},
		{root: design.MakeTree2(1, 2), pathsum: 3},
	} {
		pathsum := maxPathSum(tc.root)
		assert.Equal(t, tc.pathsum, pathsum)
	}
}
