package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestMaxFreeTime$
func TestMaxFreeTime(t *testing.T) {
	for _, tc := range []struct {
		eventTime int
		k         int
		startTime []int
		endTime   []int
		maxfree   int
	}{
		{eventTime: 5, k: 1, startTime: []int{1, 3}, endTime: []int{2, 5}, maxfree: 2},
		{eventTime: 10, k: 1, startTime: []int{0, 2, 9}, endTime: []int{1, 4, 10}, maxfree: 6},
		{eventTime: 5, k: 2, startTime: []int{0, 1, 2, 3, 4}, endTime: []int{1, 2, 3, 4, 5}, maxfree: 0},
	} {
		maxfree := maxFreeTime(tc.eventTime, tc.k, tc.startTime, tc.endTime)
		assert.Equal(t, tc.maxfree, maxfree)
	}
}
