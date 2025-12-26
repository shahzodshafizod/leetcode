package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestCountPartitions$
func TestCountPartitions(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{1, 2, 2}, count: 0},
		{nums: []int{2, 4, 6, 8}, count: 3},
		{nums: []int{1, 1, 1, 1}, count: 3},
		{nums: []int{1, 2, 3}, count: 2},
		{nums: []int{10, 20}, count: 1},
	} {
		count := countPartitions(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
