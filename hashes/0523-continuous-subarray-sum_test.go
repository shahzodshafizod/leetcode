package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCheckSubarraySum$
func TestCheckSubarraySum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		good bool
	}{
		{nums: []int{23, 2, 4, 6, 7}, k: 6, good: true},
		{nums: []int{23, 2, 6, 4, 7}, k: 6, good: true},
		{nums: []int{23, 2, 6, 4, 7}, k: 13, good: false},
		{nums: []int{1}, k: 1, good: false},
		{nums: []int{4, 1, 2, 3}, k: 6, good: true},
		{nums: []int{5, 0, 0, 0}, k: 3, good: true},
		{nums: []int{23, 2, 4, 6, 6}, k: 7, good: true},
		{nums: []int{0, 0}, k: 1, good: true},
		{nums: []int{1, 3, 6, 0, 9, 6, 9}, k: 7, good: true},
	} {
		good := checkSubarraySum(tc.nums, tc.k)
		assert.Equal(t, tc.good, good)
	}
}
