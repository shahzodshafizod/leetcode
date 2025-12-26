package trees

import (
	"sort"
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestFindMode$
func TestFindMode(t *testing.T) {
	for _, tc := range []struct {
		root  *pkg.TreeNode
		modes []int
	}{
		{root: pkg.MakeTree2(1, nil, 2, 2), modes: []int{2}},
		{root: pkg.MakeTree2(1, 1, 2), modes: []int{1}},
		{root: pkg.MakeTree2(1), modes: []int{1}},
		{root: pkg.MakeTree2(1, nil, 2, nil, 3), modes: []int{1, 2, 3}},
	} {
		modes := findMode(tc.root)
		sort.Ints(modes)
		sort.Ints(tc.modes)
		assert.Equal(t, tc.modes, modes)
	}
}
