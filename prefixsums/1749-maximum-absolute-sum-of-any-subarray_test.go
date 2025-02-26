package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestMaxAbsoluteSum$
func TestMaxAbsoluteSum(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		abssum int
	}{
		{nums: []int{1, -3, 2, 3, -4}, abssum: 5},
		{nums: []int{2, -5, 1, -4, 3, -2}, abssum: 8},
	} {
		abssum := maxAbsoluteSum(tc.nums)
		assert.Equal(t, tc.abssum, abssum)
	}
}
