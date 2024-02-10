package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestRecoverFromPreorder$
func TestRecoverFromPreorder(t *testing.T) {
	for _, tc := range []struct {
		traversal string
		root      *TreeNode
	}{
		{traversal: "1-2--3--4-5--6--7", root: makeTree(0, []any{1, 2, 5, 3, 4, 6, 7})},
		{traversal: "1-2--3---4-5--6---7", root: makeTree(0, []any{1, 2, 5, 3, nil, 6, nil, 4, nil, nil, nil, 7})},
		{traversal: "1-401--349---90--88", root: makeTree(0, []any{1, 401, nil, 349, 88, nil, nil, 90})},
	} {
		root := recoverFromPreorder(tc.traversal)
		assert.Equal(t, tc.root, root)
	}
}
