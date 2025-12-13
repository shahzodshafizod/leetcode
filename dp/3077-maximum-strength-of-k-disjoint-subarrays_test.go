package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaximumStrength$
func TestMaximumStrength(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		output int64
	}{
		{nums: []int{1, 2, 3, -1, 2}, k: 3, output: 22},
		{nums: []int{12, -2, -2, -2, -2}, k: 5, output: 64},
		{nums: []int{-1, -2, -3}, k: 1, output: -1},
		{nums: []int{-99, 85}, k: 1, output: 85},
	} {
		output := maximumStrength(tc.nums, tc.k)
		assert.Equal(t, tc.output, output)
	}
}
