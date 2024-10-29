package trees

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./trees/ -run ^TestTreeQueries$
func TestTreeQueries(t *testing.T) {
	for _, tc := range []struct {
		root    *pkg.TreeNode
		queries []int
		answer  []int
	}{
		{root: pkg.MakeTree2(5, 8, 9, 2, 1, 3, 7, 4, 6), queries: []int{3, 2, 4, 8}, answer: []int{3, 2, 3, 2}},
		{root: pkg.MakeTree2(1, nil, 5, 3, nil, 2, 4), queries: []int{3, 5, 4, 2, 4}, answer: []int{1, 0, 3, 3, 3}},
		{root: pkg.MakeTree2(1, 3, 4, 2, nil, 6, 5, nil, nil, nil, nil, nil, 7), queries: []int{4}, answer: []int{2}},
		{root: pkg.MakeTree2(2, 1, 5, nil, nil, 3, 6, nil, 4), queries: []int{1, 5, 5, 6, 4, 5}, answer: []int{3, 1, 1, 3, 2, 1}},
		{root: pkg.MakeTree2(2, nil, 5, 3, 1, nil, 4), queries: []int{1, 5, 5, 3, 4, 5}, answer: []int{3, 0, 0, 2, 2, 0}},
		{root: pkg.MakeTree2(1, 2, nil, 3, nil, 4, nil, 5, nil, 6, nil, 7, nil, 8, nil, 9, nil, 10, nil, 11, nil, 12), queries: []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, answer: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{root: pkg.MakeTree2(7, 2, 9, 1, 3, 8, 12, nil, nil, nil, 6, nil, nil, 10, nil, 4, nil, nil, 11, nil, 5), queries: []int{3, 8, 9, 10, 10, 5, 6, 12, 2, 6, 1, 11}, answer: []int{4, 5, 5, 5, 5, 4, 4, 5, 4, 4, 5, 5}},
	} {
		answer := treeQueries(tc.root, tc.queries)
		assert.Equal(t, tc.answer, answer)
	}
}
