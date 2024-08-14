package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestSmallestDistancePair$
func TestSmallestDistancePair(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		k        int
		distance int
	}{
		{nums: []int{1, 3, 1}, k: 1, distance: 0},
		{nums: []int{1, 1, 1}, k: 2, distance: 0},
		{nums: []int{1, 6, 1}, k: 3, distance: 5},
	} {
		distance := smallestDistancePair(tc.nums, tc.k)
		assert.Equal(t, tc.distance, distance)
	}
}
