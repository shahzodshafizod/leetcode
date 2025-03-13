package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestIsZeroArray$
func TestIsZeroArray(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		is      bool
	}{
		{nums: []int{1, 0, 1}, queries: [][]int{{0, 2}}, is: true},
		{nums: []int{4, 3, 2, 1}, queries: [][]int{{1, 3}, {0, 2}}, is: false},
		{nums: []int{2}, queries: [][]int{{0, 0}, {0, 0}, {0, 0}}, is: true},
	} {
		is := isZeroArray(tc.nums, tc.queries)
		assert.Equal(t, tc.is, is)
	}
}
