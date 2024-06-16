package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinPatches$
func TestMinPatches(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		n       int
		patches int
	}{
		{nums: []int{1, 3}, n: 6, patches: 1},
		{nums: []int{1, 5, 10}, n: 20, patches: 2},
		{nums: []int{1, 2, 2}, n: 5, patches: 0},
	} {
		patches := minPatches(tc.nums, tc.n)
		assert.Equal(t, tc.patches, patches)
	}
}
