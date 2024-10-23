package unionfinds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./unionfinds/ -run ^TestCanReachCorner$
func TestCanReachCorner(t *testing.T) {
	for _, tc := range []struct {
		xCorner int
		yCorner int
		circles [][]int
		can     bool
	}{
		{xCorner: 3, yCorner: 4, circles: [][]int{{2, 1, 1}}, can: true},
		{xCorner: 3, yCorner: 3, circles: [][]int{{1, 1, 2}}, can: false},
		{xCorner: 3, yCorner: 3, circles: [][]int{{2, 1, 1}, {1, 2, 1}}, can: false},
		{xCorner: 4, yCorner: 4, circles: [][]int{{5, 5, 1}}, can: true},
		{xCorner: 3, yCorner: 3, circles: [][]int{{2, 1000, 997}, {1000, 2, 997}}, can: true},
		{xCorner: 3, yCorner: 4, circles: [][]int{{5, 6, 2}}, can: true},
		{xCorner: 3, yCorner: 3, circles: [][]int{{1, 100, 2}, {3, 100, 2}}, can: true},
		{xCorner: 4, yCorner: 4, circles: [][]int{{2, 6, 2}, {6, 2, 2}, {6, 6, 2}}, can: true},
		{xCorner: 3, yCorner: 3, circles: [][]int{{4, 4, 2}}, can: false},
		{xCorner: 22742157, yCorner: 210809967, circles: [][]int{{22741186, 210810964, 200}, {22741869, 210809432, 165}, {22742256, 210810275, 182}, {22742089, 210809693, 129}, {22741912, 210810128, 196}, {22741658, 210809205, 144}, {22741648, 210809094, 118}, {22741920, 210809808, 128}}, can: false},
		{xCorner: 8, yCorner: 10, circles: [][]int{{5, 2, 2}, {6, 7, 2}, {3, 1, 1}, {7, 5, 1}, {5, 3, 1}, {3, 7, 3}, {1, 7, 1}}, can: false},
	} {
		can := canReachCorner(tc.xCorner, tc.yCorner, tc.circles)
		assert.Equal(t, tc.can, can)
	}
}
