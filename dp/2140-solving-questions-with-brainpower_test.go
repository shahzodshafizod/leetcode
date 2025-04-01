package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMostPoints$
func TestMostPoints(t *testing.T) {
	for _, tc := range []struct {
		questions [][]int
		points    int64
	}{
		{questions: [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}, points: 5},
		{questions: [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, points: 7},
	} {
		points := mostPoints(tc.questions)
		assert.Equal(t, tc.points, points)
	}
}
