package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxFreeTime$
func TestMaxFreeTime(t *testing.T) {
	for _, tc := range []struct {
		eventTime int
		startTime []int
		endTime []int
		maxfree int
	}{
		{eventTime: 5, startTime: []int{1, 3}, endTime: []int{2, 5}, maxfree: 2},
		{eventTime: 10, startTime: []int{0, 7, 9}, endTime: []int{1, 8, 10}, maxfree: 7},
		{eventTime: 10, startTime: []int{0, 3, 7, 9}, endTime: []int{1, 4, 8, 10}, maxfree: 6},
		{eventTime: 5, startTime: []int{0, 1, 2, 3, 4}, endTime: []int{1, 2, 3, 4, 5}, maxfree: 0},
		{eventTime: 34, startTime: []int{0, 17}, endTime: []int{14, 19}, maxfree: 18},
		{eventTime: 37, startTime: []int{5, 14, 27, 34}, endTime: []int{13, 18, 31, 37}, maxfree: 16},
		{eventTime: 54, startTime: []int{9, 24, 45, 50, 53}, endTime: []int{15, 26, 50, 53, 54}, maxfree: 30},
		{eventTime: 41, startTime: []int{17, 24}, endTime: []int{19, 25}, maxfree: 24},
		{eventTime: 86, startTime: []int{22, 82}, endTime: []int{66, 85}, maxfree: 38},
	} {
		maxfree := maxFreeTime(tc.eventTime, tc.startTime, tc.endTime)
		assert.Equal(t, tc.maxfree, maxfree)
	}
}
