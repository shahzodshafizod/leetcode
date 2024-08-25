package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestPostorderTraversal$
func TestPostorderTraversal(t *testing.T) {
	for _, tc := range []struct {
		vals  []any
		order []int
	}{
		{vals: []any{1, nil, 2, nil, nil, 3}, order: []int{3, 2, 1}},
		{vals: []any{}, order: []int{}},
		{vals: []any{1}, order: []int{1}},
	} {
		var root = makeTree(0, tc.vals)
		var order = postorderTraversal(root)
		assert.Equal(t, tc.order, order)
	}
}
