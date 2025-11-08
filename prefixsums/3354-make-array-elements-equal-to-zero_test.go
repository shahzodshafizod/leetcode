package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestCountValidSelections$
func TestCountValidSelections(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		count int
	}{
		{nums: []int{1, 0, 2, 0, 3}, count: 2},
		{nums: []int{2, 3, 4, 0, 4, 1, 0}, count: 0},
	} {
		count := countValidSelections(tc.nums)
		assert.Equal(t, tc.count, count)
	}
}
