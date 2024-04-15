package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestSumNumbers$
func TestSumNumbers(t *testing.T) {
	for _, tc := range []struct {
		root *TreeNode
		sum  int
	}{
		{root: makeTree(0, []any{1, 2, 3}), sum: 25},
		{root: makeTree(0, []any{4, 9, 0, 5, 1}), sum: 1026},
		{root: makeTree(0, []any{0, 1}), sum: 1},
		{root: makeTree(0, []any{4, 9, 0, nil, 1}), sum: 531},
	} {
		sum := sumNumbers(tc.root)
		assert.Equal(t, tc.sum, sum)
	}
}
