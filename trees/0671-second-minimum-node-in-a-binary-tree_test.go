package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestFindSecondMinimumValue$
func TestFindSecondMinimumValue(t *testing.T) {
	for _, tc := range []struct {
		root   *pkg.TreeNode
		second int
	}{
		{
			root:   pkg.MakeTree(0, []any{2, 2, 5, nil, nil, 5, 7}),
			second: 5,
		},
		{
			root:   pkg.MakeTree(0, []any{2, 2, 2}),
			second: -1,
		},
		{
			root:   pkg.MakeTree(0, []any{1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1}),
			second: 2,
		},
		{
			root:   pkg.MakeTree(0, []any{2, 2, 2147483647}),
			second: 2147483647,
		},
		{
			root:   pkg.MakeTree(0, []any{5, 8, 5}),
			second: 8,
		},
	} {
		second := findSecondMinimumValue(tc.root)
		assert.Equal(t, tc.second, second)
	}
}
