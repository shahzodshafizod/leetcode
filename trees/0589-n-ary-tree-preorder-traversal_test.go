package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestPreorder$
func TestPreorder(t *testing.T) {
	for _, tc := range []struct {
		root   *design.TNode
		values []int
	}{
		{
			root:   design.MakeNAryTree([]any{1, nil, 3, 2, 4, nil, 5, 6}),
			values: []int{1, 3, 5, 6, 2, 4},
		},
		{
			root:   design.MakeNAryTree([]any{1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14}),
			values: []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
		},
		{
			root:   design.MakeNAryTree([]any{1, nil}),
			values: []int{1},
		},
		{
			root:   design.MakeNAryTree([]any{}),
			values: []int{},
		},
	} {
		values := preorder(tc.root)
		assert.Equal(t, tc.values, values)
	}
}
