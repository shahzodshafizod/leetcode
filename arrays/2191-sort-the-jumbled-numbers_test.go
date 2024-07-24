package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestSortJumbled$
func TestSortJumbled(t *testing.T) {
	for _, tc := range []struct {
		mapping []int
		nums    []int
		sorted  []int
	}{
		{mapping: []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, nums: []int{991, 338, 38}, sorted: []int{338, 38, 991}},
		{mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, nums: []int{789, 456, 123}, sorted: []int{123, 456, 789}},
		{mapping: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, sorted: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{mapping: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, nums: []int{9, 99, 999, 9999, 999999999}, sorted: []int{9, 99, 999, 9999, 999999999}},
	} {
		sorted := sortJumbled(tc.mapping, tc.nums)
		assert.Equal(t, tc.sorted, sorted)
	}
}
