package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestMaxDepthNary$
func TestMaxDepthNary(t *testing.T) {
	for _, tc := range []struct {
		vals   []any
		output int
	}{
		{vals: []any{1, nil, 3, 2, 4, nil, 5, 6}, output: 3},
		{vals: []any{1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14}, output: 5},
		{vals: []any{}, output: 0},
		{vals: []any{1}, output: 1},
	} {
		root := pkg.MakeNAryTree(tc.vals)
		output := maxDepthNary(root)
		assert.Equal(t, tc.output, output)
	}
}
