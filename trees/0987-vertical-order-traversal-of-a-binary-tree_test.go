package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestVerticalTraversal$
func TestVerticalTraversal(t *testing.T) {
	for _, tc := range []struct {
		root   *pkg.TreeNode
		output [][]int
	}{
		{root: pkg.MakeTree2(3, 9, 20, nil, nil, 15, 7), output: [][]int{{9}, {3, 15}, {20}, {7}}},
		{root: pkg.MakeTree2(1, 2, 3, 4, 5, 6, 7), output: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
		{root: pkg.MakeTree2(1, 2, 3, 4, 6, 5, 7), output: [][]int{{4}, {2}, {1, 5, 6}, {3}, {7}}},
	} {
		output := verticalTraversal(tc.root)
		assert.Equal(t, tc.output, output)
	}
}
