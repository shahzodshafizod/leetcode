package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestLargestValues$
func TestLargestValues(t *testing.T) {
	for _, tc := range []struct {
		root    *pkg.TreeNode
		largest []int
	}{
		{root: pkg.MakeTree2(), largest: []int{}},
		{root: pkg.MakeTree2(1, 2, 3), largest: []int{1, 3}},
		{root: pkg.MakeTree2(95962, nil), largest: []int{95962}},
		{root: pkg.MakeTree2(1, 3, 2, 5, 3, nil, 9), largest: []int{1, 3, 9}},
		{root: pkg.MakeTree2(-1395, 32448), largest: []int{-1395, 32448}},
		{root: pkg.MakeTree2(-38429, nil, -71908, -29625, -46622, 7696, -68937, -7431, -19170, 69587, 45237), largest: []int{-38429, -71908, -29625, 7696, 69587}},
		// {root: pkg.MakeTree2(-23351, -99006, nil, -63508, 86343, -66012, 54651, -38287, 2762, 39604, nil, -33734, 2587, 66356, -99032, 50112, 37759, 59078, -89953, 44359, -73181, -18323), largest: []int{-23351, -99006, 86343, 54651, 66356, 59078}},
	} {
		largest := largestValues(tc.root)
		assert.Equal(t, tc.largest, largest)
	}
}
