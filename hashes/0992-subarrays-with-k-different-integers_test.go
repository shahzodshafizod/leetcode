package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestSubarraysWithKDistinct$
func TestSubarraysWithKDistinct(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		number int
	}{
		{nums: []int{1, 2, 1, 2, 3}, k: 2, number: 7},
		{nums: []int{1, 2, 1, 3, 4}, k: 3, number: 3},
	} {
		number := subarraysWithKDistinct(tc.nums, tc.k)
		assert.Equal(t, tc.number, number)
	}
}
