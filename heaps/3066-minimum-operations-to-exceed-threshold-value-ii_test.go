package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int
	}{
		{nums: []int{2, 11, 10, 1, 3}, k: 10, count: 2},
		{nums: []int{1, 1, 2, 4, 9}, k: 20, count: 4},
		{nums: []int{80, 47}, k: 81, count: 1},
		{nums: []int{9, 98, 52, 8}, k: 2, count: 0},
		{nums: []int{999999999, 999999999, 999999999}, k: 1000000000, count: 2},
		{nums: []int{62, 32, 62, 73, 58, 56, 68, 50}, k: 74, count: 4},
		{nums: []int{69, 89, 57, 31, 84, 97, 50, 38, 91, 86}, k: 91, count: 4},
		{nums: []int{1000000000, 999999999, 1000000000, 999999999, 1000000000, 999999999}, k: 1000000000, count: 2},
	} {
		count := minOperations(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
