package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCalculateMinimumHP$
func TestCalculateMinimumHP(t *testing.T) {
	for _, tc := range []struct {
		dungeon     [][]int
		healthPoint int
	}{
		{dungeon: [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, healthPoint: 7},
		{dungeon: [][]int{{0}}, healthPoint: 1},
		{dungeon: [][]int{{0, 0}}, healthPoint: 1},
		{dungeon: [][]int{{0, 100, -98}}, healthPoint: 1},
		{dungeon: [][]int{{100}}, healthPoint: 1},
		{dungeon: [][]int{{10}}, healthPoint: 1},
		{dungeon: [][]int{{-3, 5}}, healthPoint: 4},
		{dungeon: [][]int{{-10}}, healthPoint: 11},
		{dungeon: [][]int{{-2, -3, 3, -5, -10}}, healthPoint: 18},
		{dungeon: [][]int{{2, 3, 3, 5, 10}}, healthPoint: 1},
	} {
		healthPoint := calculateMinimumHP(tc.dungeon)
		assert.Equal(t, tc.healthPoint, healthPoint)
	}
}
