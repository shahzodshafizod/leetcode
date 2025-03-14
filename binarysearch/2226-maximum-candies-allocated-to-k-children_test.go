package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMaximumCandies$
func TestMaximumCandies(t *testing.T) {
	for _, tc := range []struct {
		candies []int
		k       int64
		count   int
	}{
		{candies: []int{5, 8, 6}, k: 3, count: 5},
		{candies: []int{2, 5}, k: 11, count: 0},
		{candies: []int{4, 7, 5}, k: 4, count: 3},
		{candies: []int{1, 2, 3, 4, 10}, k: 5, count: 3},
		{candies: []int{5, 6, 4, 10, 10, 1, 1, 2, 2, 2}, k: 9, count: 3},
		{candies: []int{1, 2, 6, 8, 6, 7, 3, 5, 2, 5}, k: 3, count: 6},
		{candies: []int{4, 9, 4, 7, 8, 10, 3, 10, 3, 9}, k: 9, count: 4},
		{candies: []int{9, 10, 1, 2, 10, 9, 9, 10, 2, 2}, k: 3, count: 10},
	} {
		count := maximumCandies(tc.candies, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
