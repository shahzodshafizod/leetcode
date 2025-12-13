package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountSubarraysWithAndK$
func TestCountSubarraysWithAndK(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		output int64
	}{
		{nums: []int{1, 1, 1}, k: 1, output: 6}, // All subarrays
		{nums: []int{1, 1, 2}, k: 1, output: 3}, // [1], [1], [1,1]
		{nums: []int{1, 2, 3}, k: 2, output: 2}, // [2], [1,2]
		{nums: []int{4, 5, 6}, k: 4, output: 4}, // [4], [4,5], [5,6], [4,5,6]
		{nums: []int{7, 7, 7}, k: 7, output: 6}, // All subarrays
	} {
		output := countSubarraysWithAndK(tc.nums, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
