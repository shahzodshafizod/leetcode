package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestCountInterestingSubarrays$
func TestCountInterestingSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		modulo int
		k      int
		count  int64
	}{
		{nums: []int{3, 2, 4}, modulo: 2, k: 1, count: 3},
		{nums: []int{3, 1, 9, 6}, modulo: 3, k: 0, count: 2},
	} {
		count := countInterestingSubarrays(tc.nums, tc.modulo, tc.k)
		assert.Equal(t, tc.count, count)
	}
}
