package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestLevelOrder429$
func TestLevelOrder429(t *testing.T) {
	for _, tc := range []struct {
		root   *design.TNode
		values [][]int
	}{
		{
			root:   design.MakeNAryTree([]any{1, nil, 3, 2, 4, nil, 5, 6}),
			values: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			root:   design.MakeNAryTree([]any{1, nil, 2, 3, 4, 5, nil, nil, 6, 7, nil, 8, nil, 9, 10, nil, nil, 11, nil, 12, nil, 13, nil, nil, 14}),
			values: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	} {
		values := levelOrder429(tc.root)
		assert.Equal(t, tc.values, values)
	}
}
