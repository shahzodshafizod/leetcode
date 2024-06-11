package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestRelativeSortArray$
func TestRelativeSortArray(t *testing.T) {
	for _, tc := range []struct {
		arr1   []int
		arr2   []int
		sorted []int
	}{
		{arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, arr2: []int{2, 1, 4, 3, 9, 6}, sorted: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}},
		{arr1: []int{28, 6, 22, 8, 44, 17}, arr2: []int{22, 28, 8, 6}, sorted: []int{22, 28, 8, 6, 17, 44}},
	} {
		sorted := relativeSortArray(tc.arr1, tc.arr2)
		assert.Equal(t, tc.sorted, sorted)
	}
}
