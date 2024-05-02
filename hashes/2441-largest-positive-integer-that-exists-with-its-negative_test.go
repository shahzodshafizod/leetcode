package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindMaxK$
func TestFindMaxK(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
	}{
		{nums: []int{-1, 2, -3, 3}, k: 3},
		{nums: []int{-1, 10, 6, 7, -7, 1}, k: 7},
		{nums: []int{-10, 8, 6, 7, -2, -3}, k: -1},
	} {
		k := findMaxK(tc.nums)
		assert.Equal(t, tc.k, k)
	}
}
