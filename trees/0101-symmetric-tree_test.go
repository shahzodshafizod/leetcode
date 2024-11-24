package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsSymmetric$
func TestIsSymmetric(t *testing.T) {
	for _, tc := range []struct {
		root   *pkg.TreeNode
		mirror bool
	}{
		{root: pkg.MakeTree2(1, 2, 2, 3, 4, 4, 3), mirror: true},
		{root: pkg.MakeTree2(1, 2, 2, nil, 3, nil, 3), mirror: false},
		{root: pkg.MakeTree2(1, 2, 2, 5, nil, nil, 5, 6, nil, nil, 6), mirror: true},
		{root: pkg.MakeTree2(1, 2, 2, 2, nil, 2), mirror: false},
		{root: pkg.MakeTree2(1, 2, 3), mirror: false},
		{root: pkg.MakeTree2(1), mirror: true},
		{root: pkg.MakeTree2(2, 3, 3, 4, 5, 5, 4, nil, nil, 8, 9, 9, 8), mirror: true},
		{root: pkg.MakeTree2(1, 2, 2, nil, 3, 2), mirror: false},
	} {
		mirror := isSymmetric(tc.root)
		assert.Equal(t, tc.mirror, mirror)
	}
}
