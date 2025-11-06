package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMaxFrequency$
func TestMaxFrequency(t *testing.T) {
	for _, tc := range []struct {
		nums          []int
		k             int
		numOperations int
		freq          int
	}{
		{nums: []int{1, 4, 5}, k: 1, numOperations: 2, freq: 2},
		{nums: []int{5, 11, 20, 20}, k: 5, numOperations: 1, freq: 2},
	} {
		freq := maxFrequency(tc.nums, tc.k, tc.numOperations)
		assert.Equal(t, tc.freq, freq)
	}
}
