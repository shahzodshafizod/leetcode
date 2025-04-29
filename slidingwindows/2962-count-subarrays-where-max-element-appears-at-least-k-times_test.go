package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestCountSubarrays2962$
func TestCountSubarrays2962(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int64
	}{
		{nums: []int{10}, k: 1, count: 1},
		{nums: []int{1, 4, 2, 1}, k: 3, count: 0},
		{nums: []int{1, 1, 1, 1, 1}, k: 5, count: 1},
		{nums: []int{1, 3, 2, 3, 3}, k: 2, count: 6},
		{nums: []int{3, 3, 3, 3, 3}, k: 3, count: 6},
		{nums: []int{1, 2, 3, 4, 5}, k: 1, count: 5},
		{nums: []int{5, 5, 1, 5, 2, 5, 5}, k: 4, count: 3},
		{nums: []int{1, 2, 2, 2, 3, 2, 2}, k: 3, count: 0},
		{nums: []int{4, 4, 4, 1, 4, 4, 4, 4}, k: 5, count: 6},
		{nums: []int{1, 5, 5, 5, 2, 5, 5, 1, 5}, k: 5, count: 7},
		{nums: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, k: 8, count: 6},
		{nums: []int{1000000, 1, 1000000, 2, 1000000}, k: 2, count: 5},
		// {nums: []int{28, 5, 58, 91, 24, 91, 53, 9, 48, 85, 16, 70, 91, 91, 47, 91, 61, 4, 54, 61, 49}, k: 1, count: 187},
		// {nums: []int{1, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}, k: 2, count: 1221},
		// {nums: []int{1, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}, k: 1, count: 1273},
	} {
		count := countSubarrays2962(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
