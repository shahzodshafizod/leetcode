package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestArrayRankTransform$
func TestArrayRankTransform(t *testing.T) {
	for _, tc := range []struct {
		arr   []int
		ranks []int
	}{
		{arr: []int{40, 10, 20, 30}, ranks: []int{4, 1, 2, 3}},
		{arr: []int{100, 100, 100}, ranks: []int{1, 1, 1}},
		{arr: []int{37, 12, 28, 9, 100, 56, 80, 5, 12}, ranks: []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
		{arr: []int{40, 10, 10, 20, 30}, ranks: []int{4, 1, 1, 2, 3}},
		{arr: []int{}, ranks: []int{}},
		{arr: []int{-1000000000, -1000000000, -1000000000, 1000000000, 1000000000, 1000000000}, ranks: []int{1, 1, 1, 2, 2, 2}},
		{arr: []int{0, 0, 0, -1}, ranks: []int{2, 2, 2, 1}},
	} {
		ranks := arrayRankTransform(tc.arr)
		assert.Equal(t, tc.ranks, ranks)
	}
}
