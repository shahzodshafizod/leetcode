package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestFindPoisonedDuration$
func TestFindPoisonedDuration(t *testing.T) {
	for _, tc := range []struct {
		timeSeries []int
		duration   int
		total      int
	}{
		{timeSeries: []int{1, 4}, duration: 2, total: 4},
		{timeSeries: []int{1, 2}, duration: 2, total: 3},
		{timeSeries: []int{1, 2, 3, 4, 5}, duration: 5, total: 9},
		{timeSeries: []int{1}, duration: 10, total: 10},
		{timeSeries: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, duration: 1, total: 9},
	} {
		total := findPoisonedDuration(tc.timeSeries, tc.duration)
		assert.Equal(t, tc.total, total)
	}
}
