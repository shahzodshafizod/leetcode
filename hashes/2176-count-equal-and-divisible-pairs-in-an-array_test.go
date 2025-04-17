package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCountPairs$
func TestCountPairs(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int
	}{
		{nums: []int{3, 1, 2, 2, 2, 1, 3}, k: 2, count: 4},
		{nums: []int{1, 2, 3, 4}, k: 1, count: 0},
	} {
		count := countPairs(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
