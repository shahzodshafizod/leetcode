package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestUniquePathsWithObstacles$
func TestUniquePathsWithObstacles(t *testing.T) {
	for _, tc := range []struct {
		obstacleGrid [][]int
		pathCount    int
	}{
		{obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, pathCount: 2},
		{obstacleGrid: [][]int{{0, 1}, {0, 0}}, pathCount: 1},
		{obstacleGrid: [][]int{{1, 0}, {0, 0}}, pathCount: 0},
		{obstacleGrid: [][]int{{1}}, pathCount: 0},
		{obstacleGrid: [][]int{{0}}, pathCount: 1},
		{obstacleGrid: [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}, pathCount: 7},
		{obstacleGrid: [][]int{{0, 0}, {0, 1}}, pathCount: 0},
		{obstacleGrid: [][]int{{0, 0}, {1, 1}, {0, 0}}, pathCount: 0},
		{obstacleGrid: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 1}, {0, 0, 0, 1, 0}, {0, 0, 1, 0, 0}}, pathCount: 0},
	} {
		pathCount := uniquePathsWithObstacles(tc.obstacleGrid)
		assert.Equal(t, tc.pathCount, pathCount)
	}
}
