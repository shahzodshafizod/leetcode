package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxScoreSightseeingPair$
func TestMaxScoreSightseeingPair(t *testing.T) {
	for _, tc := range []struct {
		values []int
		score  int
	}{
		{values: []int{8, 1, 5, 2, 6}, score: 11},
		{values: []int{1, 2}, score: 2},
	} {
		score := maxScoreSightseeingPair(tc.values)
		assert.Equal(t, tc.score, score)
	}
}
