package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestCountSubarrays2302$
func TestCountSubarrays2302(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int64
		count int64
	}{
		{nums: []int{2, 1, 4, 3, 5}, k: 10, count: 6},
		{nums: []int{1, 1, 1}, k: 5, count: 5},
	} {
		count := countSubarrays2302(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
