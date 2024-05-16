package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestEvaluateTree$
func TestEvaluateTree(t *testing.T) {
	for _, tc := range []struct {
		root   *TreeNode
		result bool
	}{
		{root: makeTree(0, []any{2, 1, 3, nil, nil, 0, 1}), result: true},
		{root: makeTree(0, []any{0}), result: false},
	} {
		result := evaluateTree(tc.root)
		assert.Equal(t, tc.result, result)
	}
}
