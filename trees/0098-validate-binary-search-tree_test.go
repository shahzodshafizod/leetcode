package trees

import (
	"testing"

	"github.com/shahzodshafizod/alkhwarizmi/design"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsValidBST$
func TestIsValidBST(t *testing.T) {
	for _, tc := range []struct {
		root  *design.TreeNode
		valid bool
	}{
		{root: design.MakeTree2(2, 1, 3), valid: true},
		{root: design.MakeTree2(5, 1, 4, nil, nil, 3, 6), valid: false},
		{root: design.MakeTree2(10, 5, 15, nil, nil, 6, 20), valid: false},
		{root: design.MakeTree2(12, 8, 18, 5, 10, 14, 25), valid: true},
		{root: design.MakeTree2(16, 8, 22, 9, nil, 19, 25), valid: false},
		{root: design.MakeTree2(13, 6, 17, 2, nil, 10, 22), valid: false},
		{root: design.MakeTree2(), valid: true},
		{root: design.MakeTree2(12, 7, 18, 5, 9, 16, 25), valid: true},
		{root: design.MakeTree2(1), valid: true},
	} {
		valid := isValidBST(tc.root)
		assert.Equal(t, tc.valid, valid)
	}
}
