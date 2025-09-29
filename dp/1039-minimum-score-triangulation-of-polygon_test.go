package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinScoreTriangulation$
func TestMinScoreTriangulation(t *testing.T) {
	for _, tc := range []struct {
		values []int
		score  int
	}{
		{values: []int{1, 2, 3}, score: 6},
		{values: []int{3, 7, 4, 5}, score: 144},
		{values: []int{1, 3, 1, 4, 1, 5}, score: 13},
		{values: []int{4, 3, 4, 3, 5}, score: 132},
		{values: []int{2, 1, 4, 4}, score: 24},
		{values: []int{1, 2, 3, 4}, score: 18},
		// {values: []int{35, 73, 90, 27, 71, 80, 21, 33, 33, 13, 48, 12, 68, 70, 80, 36, 66, 3, 70, 58}, score: 140295},
	} {
		score := minScoreTriangulation(tc.values)
		assert.Equal(t, tc.score, score)
	}
}
