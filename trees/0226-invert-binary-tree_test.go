package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestInvertTree$
func TestInvertTree(t *testing.T) {
	for _, tc := range []struct {
		root   *pkg.TreeNode
		output *pkg.TreeNode
	}{
		{root: pkg.MakeTree2(4, 2, 7, 1, 3, 6, 9), output: pkg.MakeTree2(4, 7, 2, 9, 6, 3, 1)},
		{root: pkg.MakeTree2(2, 1, 3), output: pkg.MakeTree2(2, 3, 1)},
		{root: pkg.MakeTree2(), output: pkg.MakeTree2()},
		{root: pkg.MakeTree2(1), output: pkg.MakeTree2(1)},
	} {
		output := invertTree(tc.root)
		assert.Equal(t, tc.output, output)
	}
}
