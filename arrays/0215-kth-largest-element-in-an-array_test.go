package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindKthLargest$
func TestFindKthLargest(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		large int
	}{
		{nums: []int{5, 3, 1, 6, 4, 2}, k: 2, large: 5},
		{nums: []int{5, 3, 1, 6, 4, 4, 2}, k: 2, large: 5},
		{nums: []int{2, 3, 1, 2, 4, 2}, k: 4, large: 2},
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 2, large: 5},
		{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, large: 4},
		{nums: []int{3, 2, 3, 1, 2, 1, 4, 5, 5, 6}, k: 4, large: 4},
	} {
		large := findKthLargest(tc.nums, tc.k)
		assert.Equal(t, tc.large, large)
	}
}
