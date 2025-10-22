package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaximumLength3177$
func TestMaximumLength3177(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		length int
	}{
		{nums: []int{1, 2, 1, 1, 3}, k: 2, length: 4},
		{nums: []int{1, 2, 3, 4, 5, 1}, k: 0, length: 2},
	} {
		length := maximumLength3177(tc.nums, tc.k)
		assert.Equal(t, tc.length, length)
	}
}
