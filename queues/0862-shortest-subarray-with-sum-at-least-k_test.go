package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestShortestSubarray$
func TestShortestSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		length int
	}{
		{nums: []int{1}, k: 1, length: 1},
		{nums: []int{1, 2}, k: 4, length: -1},
		{nums: []int{2, -1, 2}, k: 3, length: 3},
		{nums: []int{48, 99, 37, 4, -31}, k: 140, length: 2},
		{nums: []int{44, -25, 75, -50, -38, -42, -32, -6, -40, -46, -47}, k: 19, length: 1},
		{nums: []int{10, -5, 200}, k: 150, length: 1},
		{nums: []int{2, -1, 1, 3}, k: 4, length: 2},
	} {
		length := shortestSubarray(tc.nums, tc.k)
		assert.Equal(t, tc.length, length)
	}
}
