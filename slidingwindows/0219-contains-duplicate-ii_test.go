package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestContainsNearbyDuplicate$
func TestContainsNearbyDuplicate(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		k        int
		contains bool
	}{
		{nums: []int{1, 2, 3, 1}, k: 3, contains: true},
		{nums: []int{1, 0, 1, 1}, k: 1, contains: true},
		{nums: []int{1, 2, 3, 1, 2, 3}, k: 2, contains: false},
		{nums: []int{1, 2, 3, 2, 3, 3}, k: 3, contains: true},
	} {
		contains := containsNearbyDuplicate(tc.nums, tc.k)
		assert.Equal(t, tc.contains, contains)
	}
}
