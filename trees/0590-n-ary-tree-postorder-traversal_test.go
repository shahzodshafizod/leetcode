package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestPostorder$
func TestPostorder(t *testing.T) {
	for _, tc := range []struct {
		root  *Node
		order []int
	}{
		{root: makeNAryTree([]any{1, nil, 3, 2, 4, nil, 5, 6}), order: []int{5, 6, 3, 2, 4, 1}},
		{root: makeNAryTree([]any{1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14}), order: []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1}},
		{root: makeNAryTree([]any{}), order: []int{}},
	} {
		order := postorder(tc.root)
		assert.Equal(t, tc.order, order)
	}
}
