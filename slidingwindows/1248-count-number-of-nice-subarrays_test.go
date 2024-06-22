package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestNumberOfSubarrays$
func TestNumberOfSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		k         int
		subarrays int
	}{
		{nums: []int{1, 1, 2, 1, 1}, k: 3, subarrays: 2},
		{nums: []int{2, 4, 6}, k: 1, subarrays: 0},
		{nums: []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, k: 2, subarrays: 16},
		{nums: []int{0, 0, 0, 1, 1, 1, 1, 0, 0}, k: 4, subarrays: 12},
		{nums: []int{4, 1, 2, 3, 7, 9, 1, 2, 8}, k: 2, subarrays: 8},
	} {
		subarrays := numberOfSubarrays(tc.nums, tc.k)
		assert.Equal(t, tc.subarrays, subarrays)
	}
}
