package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestGetFinalState$
func TestGetFinalState(t *testing.T) {
	for _, tc := range []struct {
		nums       []int
		k          int
		multiplier int
		final      []int
	}{
		{nums: []int{2, 1}, k: 1, multiplier: 2, final: []int{2, 2}},
		{nums: []int{1, 2}, k: 3, multiplier: 4, final: []int{16, 8}},
		{nums: []int{2, 1, 3, 5, 6}, k: 5, multiplier: 2, final: []int{8, 4, 6, 5, 6}},
	} {
		final := getFinalState(tc.nums, tc.k, tc.multiplier)
		assert.Equal(t, tc.final, final)
	}
}
