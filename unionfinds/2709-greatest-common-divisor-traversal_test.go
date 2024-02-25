package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestCanTraverseAllPairs$
func TestCanTraverseAllPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		can  bool
	}{
		{nums: []int{2, 3, 6}, can: true},
		{nums: []int{3, 9, 5}, can: false},
		{nums: []int{4, 3, 12, 8}, can: true},
		{nums: []int{12, 2}, can: true},
		{nums: []int{1}, can: true},
		{nums: []int{1, 1}, can: false},
		{nums: []int{30, 30}, can: true},
	} {
		can := canTraverseAllPairs(tc.nums)
		assert.Equal(t, tc.can, can)
	}
}
