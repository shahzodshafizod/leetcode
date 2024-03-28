package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMaxSubarrayLength$
func TestMaxSubarrayLength(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		length int
	}{
		{nums: []int{1, 2, 3, 1, 2, 3, 1, 2}, k: 2, length: 6},
		{nums: []int{1, 2, 3, 1, 2, 3, 2, 1}, k: 2, length: 6},
		{nums: []int{1, 2, 3, 1, 2, 3}, k: 2, length: 6},
		{nums: []int{1, 2, 1, 2, 1, 2, 1, 2}, k: 1, length: 2},
		{nums: []int{5, 5, 5, 5, 5, 5, 5}, k: 4, length: 4},
		{nums: []int{1, 11}, k: 3, length: 2},
		{nums: []int{2}, k: 2, length: 1},
	} {
		length := maxSubarrayLength(tc.nums, tc.k)
		assert.Equal(t, tc.length, length)
	}
}
