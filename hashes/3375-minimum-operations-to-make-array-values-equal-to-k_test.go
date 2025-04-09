package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		opers int
	}{
		{nums: []int{5, 2, 5, 4, 5}, k: 2, opers: 2},
		{nums: []int{2, 1, 2}, k: 2, opers: -1},
		{nums: []int{9, 7, 5, 3}, k: 1, opers: 4},
	} {
		opers := minOperations(tc.nums, tc.k)
		assert.Equal(t, tc.opers, opers)
	}
}
