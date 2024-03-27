package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCountSubarrays$
func TestCountSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int64
		count int64
	}{
		{nums: []int{2, 1, 4, 3, 5}, k: 10, count: 6},
		{nums: []int{1, 1, 1}, k: 5, count: 5},
	} {
		count := countSubarrays(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
