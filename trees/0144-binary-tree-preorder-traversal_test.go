package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestPreorderTraversal$
func TestPreorderTraversal(t *testing.T) {
	for _, tc := range []struct {
		rootVals []any
		vals     []int
	}{
		{rootVals: []any{1, nil, 2, 3}, vals: []int{1, 2, 3}},
		{rootVals: []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}, vals: []int{1, 2, 4, 5, 6, 7, 3, 8, 9}},
		{rootVals: []any{}, vals: []int{}},
		{rootVals: []any{1}, vals: []int{1}},
	} {
		vals := preorderTraversal(pkg.MakeTree2(tc.rootVals...))
		assert.Equal(t, tc.vals, vals)
	}
}
