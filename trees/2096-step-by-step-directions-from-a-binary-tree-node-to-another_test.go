package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestGetDirections$
func TestGetDirections(t *testing.T) {
	for _, tc := range []struct {
		root       *design.TreeNode
		startValue int
		destValue  int
		directions string
	}{
		{root: design.MakeTree(0, []any{5, 1, 2, 3, nil, 6, 4}), startValue: 3, destValue: 6, directions: "UURL"},
		{root: design.MakeTree(0, []any{2, 1}), startValue: 2, destValue: 1, directions: "L"},
	} {
		directions := getDirections(tc.root, tc.startValue, tc.destValue)
		assert.Equal(t, tc.directions, directions)
	}
}
