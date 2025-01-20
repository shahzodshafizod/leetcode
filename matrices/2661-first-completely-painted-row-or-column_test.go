package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestFirstCompleteIndex$
func TestFirstCompleteIndex(t *testing.T) {
	for _, tc := range []struct {
		arr []int
		mat [][]int
		idx int
	}{
		{arr: []int{1, 3, 4, 2}, mat: [][]int{{1, 4}, {2, 3}}, idx: 2},
		{arr: []int{2, 8, 7, 4, 1, 3, 5, 6, 9}, mat: [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}, idx: 3},
	} {
		idx := firstCompleteIndex(tc.arr, tc.mat)
		assert.Equal(t, tc.idx, idx)
	}
}
