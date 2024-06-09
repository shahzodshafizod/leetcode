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
	} {
		subarrays := subarraysDivByK(tc.nums, tc.k)
		assert.Equal(t, tc.subarrays, subarrays)
	}
}
