package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestIsBalanced$
func TestIsBalanced(t *testing.T) {
	for _, tc := range []struct {
		rootVals []any
		balanced bool
	}{
		{rootVals: []any{3, 9, 20, nil, nil, 15, 7}, balanced: true},
		{rootVals: []any{1, 2, 2, 3, 3, nil, nil, 4, 4}, balanced: false},
		{rootVals: []any{}, balanced: true},
	} {
		balanced := isBalanced(pkg.MakeTree2(tc.rootVals...))
		assert.Equal(t, tc.balanced, balanced)
	}
}
