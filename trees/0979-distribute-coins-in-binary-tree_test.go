package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestDistributeCoins$
func TestDistributeCoins(t *testing.T) {
	for _, tc := range []struct {
		root  *pkg.TreeNode
		moves int
	}{
		{root: pkg.MakeTree(0, []any{3, 0, 0}), moves: 2},
		{root: pkg.MakeTree(0, []any{0, 3, 0}), moves: 3},
	} {
		moves := distributeCoins(tc.root)
		assert.Equal(t, tc.moves, moves)
	}
}
