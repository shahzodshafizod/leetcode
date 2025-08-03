package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestNumSubarrayProductLessThanK$
func TestNumSubarrayProductLessThanK(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int
	}{
		{nums: []int{10, 5, 2, 6}, k: 100, count: 8},
		{nums: []int{1, 2, 3}, k: 0, count: 0},
		{nums: []int{1, 1}, k: 1, count: 0},
		{nums: []int{1, 2, 3, 4, 5}, k: 1, count: 0},
		{nums: []int{2, 5, 10, 6}, k: 30, count: 5},
	} {
		count := numSubarrayProductLessThanK(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
