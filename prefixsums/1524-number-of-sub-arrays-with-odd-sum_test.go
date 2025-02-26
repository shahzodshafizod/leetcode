package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestNumOfSubarrays$
func TestNumOfSubarrays(t *testing.T) {
	for _, tc := range []struct {
		arr   []int
		count int
	}{
		{arr: []int{1, 3, 5}, count: 4},
		{arr: []int{2, 4, 6}, count: 0},
		{arr: []int{1, 2, 3, 4, 5, 6, 7}, count: 16},
	} {
		count := numOfSubarrays(tc.arr)
		assert.Equal(t, tc.count, count)
	}
}
