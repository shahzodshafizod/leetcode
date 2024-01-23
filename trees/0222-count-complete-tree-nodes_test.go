package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestCountNodes$
func TestCountNodes(t *testing.T) {
	for _, tc := range []struct {
		root  *TreeNode
		count int
	}{
		{root: makeTree(0, []any{1, 2, 3, 4, 5, 6, nil}), count: 6},
		{root: makeTree(0, []any{}), count: 0},
		{root: makeTree(0, []any{1}), count: 1},
	} {
		count := countNodes(tc.root)
		assert.Equal(t, tc.count, count)
	}
}
