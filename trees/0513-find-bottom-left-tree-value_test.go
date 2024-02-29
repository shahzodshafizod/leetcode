package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestFindBottomLeftValue$
func TestFindBottomLeftValue(t *testing.T) {
	for _, tc := range []struct {
		root  *TreeNode
		value int
	}{
		{root: makeTree(0, []any{2, 1, 3}), value: 1},
		{root: makeTree(0, []any{1, 2, 3, 4, nil, 5, 6, nil, nil, nil, nil, 7}), value: 7},
		{root: makeTree(0, []any{0, nil, -1}), value: -1},
	} {
		value := findBottomLeftValue(tc.root)
		assert.Equal(t, tc.value, value)
	}
}