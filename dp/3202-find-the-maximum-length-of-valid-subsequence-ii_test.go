package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaximumLength3202$
func TestMaximumLength3202(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		k      int
		length int
	}{
		{nums: []int{1, 2, 3, 4, 5}, k: 2, length: 5},
		{nums: []int{1, 4, 2, 3, 1, 4}, k: 3, length: 4},
	} {
		length := maximumLength3202(tc.nums, tc.k)
		assert.Equal(t, tc.length, length)
	}
}
