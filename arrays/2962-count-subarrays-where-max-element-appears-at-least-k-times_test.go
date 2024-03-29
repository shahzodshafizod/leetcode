package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountSubarrays2962$
func TestCountSubarrays2962(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int64
	}{
		{nums: []int{1, 3, 2, 3, 3}, k: 2, count: 6},
		{nums: []int{1, 4, 2, 1}, k: 3, count: 0},
		{nums: []int{28, 5, 58, 91, 24, 91, 53, 9, 48, 85, 16, 70, 91, 91, 47, 91, 61, 4, 54, 61, 49}, k: 1, count: 187},
	} {
		count := countSubarrays2962(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
