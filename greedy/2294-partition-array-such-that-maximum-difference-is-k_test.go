package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestPartitionArray$
func TestPartitionArray(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int
	}{
		{nums: []int{3, 6, 1, 2, 5}, k: 2, count: 2},
		{nums: []int{1, 2, 3}, k: 1, count: 2},
		{nums: []int{2, 2, 4, 5}, k: 0, count: 3},
	} {
		count := partitionArray(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
