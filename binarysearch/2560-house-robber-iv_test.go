package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMinCapability$
func TestMinCapability(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		amount int
	}{
		{nums: []int{2, 3, 5, 9}, k: 2, amount: 5},
		{nums: []int{2, 7, 9, 3, 1}, k: 2, amount: 2},
	} {
		amount := minCapability(tc.nums, tc.k)
		assert.Equal(t, tc.amount, amount)
	}
}
