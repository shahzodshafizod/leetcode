package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestSubarraySum$
func TestSubarraySum(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		k     int
		count int
	}{
		{nums: []int{1, 1, 1}, k: 2, count: 2},
		{nums: []int{1, 2, 3}, k: 3, count: 2},
	} {
		count := subarraySum(tc.nums, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
