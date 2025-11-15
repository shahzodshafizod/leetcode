package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMakeArrayIncreasing$
func TestMakeArrayIncreasing(t *testing.T) {
	for _, tc := range []struct {
		arr1 []int
		arr2 []int
		ops  int
	}{
		{arr1: []int{1, 5, 3, 6, 7}, arr2: []int{1, 3, 2, 4}, ops: 1},
		{arr1: []int{1, 5, 3, 6, 7}, arr2: []int{4, 3, 1}, ops: 2},
		{arr1: []int{1, 5, 3, 6, 7}, arr2: []int{1, 6, 3, 3}, ops: -1},
		{arr1: []int{31, 18, 1, 12, 23, 14, 25, 4, 17, 18, 29, 28, 35, 34, 19, 8, 25, 6, 35}, arr2: []int{13, 10, 25, 18, 3, 8, 37, 20, 23, 12, 9, 36, 17, 22, 29, 6, 1, 12, 37, 6, 3, 14, 37, 2}, ops: 19},
		{arr1: []int{0, 11, 6, 1, 4, 3}, arr2: []int{5, 4, 11, 10, 1, 0}, ops: 5},
	} {
		ops := makeArrayIncreasing(tc.arr1, tc.arr2)
		assert.Equal(t, tc.ops, ops)
	}
}
