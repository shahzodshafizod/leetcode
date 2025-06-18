package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestDivideArray$
func TestDivideArray(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		parts [][]int
	}{
		{nums: []int{1, 3, 4, 8, 7, 9, 3, 5, 1}, k: 2, parts: [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}}},
		{nums: []int{2, 4, 2, 2, 5, 2}, k: 2, parts: nil},
		{nums: []int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11}, k: 14, parts: [][]int{{2, 2, 2}, {4, 5, 5}, {5, 5, 7}, {7, 8, 8}, {9, 9, 10}, {11, 12, 12}}},
	} {
		parts := divideArray(tc.nums, tc.k)
		assert.Equal(t, tc.parts, parts)
	}
}
