package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestSubarraysDivByK$
func TestSubarraysDivByK(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		k         int
		subarrays int
	}{
		{nums: []int{4, 5, 0, -2, -3, 1}, k: 5, subarrays: 7},
		{nums: []int{4, 5, 0, -10, -3, 1}, k: 5, subarrays: 6},
		{nums: []int{5}, k: 9, subarrays: 0},
		{nums: []int{23, 2, 6, 4, 7}, k: 6, subarrays: 3},
		{nums: []int{23, 2, 6, 4, 7}, k: 13, subarrays: 0},
		{nums: []int{-5}, k: 5, subarrays: 1},
		{nums: []int{0}, k: 2, subarrays: 1},
	} {
		subarrays := subarraysDivByK(tc.nums, tc.k)
		assert.Equal(t, tc.subarrays, subarrays)
	}
}
