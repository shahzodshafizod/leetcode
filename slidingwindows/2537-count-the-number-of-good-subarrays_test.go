package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestCountGood$
func TestCountGood(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int64
	}{
		{nums: []int{1, 1, 1, 1, 1}, k: 10, count: 1},
		{nums: []int{3, 1, 4, 3, 2, 2, 4}, k: 2, count: 4},
	} {
		count := countGood(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
