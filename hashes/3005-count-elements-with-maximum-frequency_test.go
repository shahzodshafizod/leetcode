package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestMaxFrequencyElements$
func TestMaxFrequencyElements(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		total int
	}{
		{nums: []int{1, 2, 2, 3, 1, 4}, total: 4},
		{nums: []int{1, 2, 3, 4, 5}, total: 5},
	} {
		total := maxFrequencyElements(tc.nums)
		assert.Equal(t, tc.total, total)
	}
}
