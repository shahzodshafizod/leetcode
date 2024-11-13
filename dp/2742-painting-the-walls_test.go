package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestPaintWalls$
func TestPaintWalls(t *testing.T) {
	for _, tc := range []struct {
		cost    []int
		time    []int
		mincost int
	}{
		{cost: []int{2, 2}, time: []int{1, 1}, mincost: 2},
		{cost: []int{1, 2, 3, 2}, time: []int{1, 2, 3, 2}, mincost: 3},
		{cost: []int{2, 3, 4, 2}, time: []int{1, 1, 1, 1}, mincost: 4},
		{cost: []int{8, 7, 5, 15}, time: []int{1, 1, 2, 1}, mincost: 12},
		{cost: []int{49, 35, 32, 20, 30, 12, 42}, time: []int{1, 1, 2, 2, 1, 1, 2}, mincost: 62},
		{cost: []int{42, 8, 28, 35, 21, 13, 21, 35}, time: []int{2, 1, 1, 1, 2, 1, 1, 2}, mincost: 63},
		{cost: []int{26, 53, 10, 24, 25, 20, 63, 51}, time: []int{1, 1, 1, 1, 2, 2, 2, 1}, mincost: 55},
		{cost: []int{76, 25, 96, 46, 85, 19, 29, 88, 2, 5}, time: []int{1, 2, 1, 3, 1, 3, 3, 3, 2, 1}, mincost: 46},
	} {
		mincost := paintWalls(tc.cost, tc.time)
		assert.Equal(t, tc.mincost, mincost)
	}
}
