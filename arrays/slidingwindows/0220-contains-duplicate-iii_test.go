package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/slidingwindows/ -run ^TestContainsNearbyAlmostDuplicate$
func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		indexDiff int
		valueDiff int
		contains  bool
	}{
		{nums: []int{1, 2, 3, 1}, indexDiff: 3, valueDiff: 0, contains: true},
		{nums: []int{1, 5, 9, 1, 5, 9}, indexDiff: 2, valueDiff: 3, contains: false},
		{nums: []int{8, 7, 15, 1, 6, 1, 9, 15}, indexDiff: 1, valueDiff: 3, contains: true},
		{nums: []int{1, 0, 1, 1}, indexDiff: 1, valueDiff: 2, contains: true},
		{nums: []int{1, 2}, indexDiff: 1, valueDiff: -1, contains: false},
		{nums: []int{2, 3}, indexDiff: 1, valueDiff: -1, contains: false},
		{nums: []int{-2147483648 - 2147483647}, indexDiff: 1, valueDiff: 1, contains: false},
		{nums: []int{7, 2, 3}, indexDiff: 2, valueDiff: 3, contains: true},
		{nums: []int{-1, -1}, indexDiff: 1, valueDiff: -1, contains: false},
		{nums: []int{2, 2}, indexDiff: 2, valueDiff: 0, contains: true},
	} {
		contains := containsNearbyAlmostDuplicate(tc.nums, tc.indexDiff, tc.valueDiff)
		assert.Equal(t, tc.contains, contains)
	}
}
