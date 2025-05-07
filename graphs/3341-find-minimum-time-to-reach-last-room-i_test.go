package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestMinTimeToReach$
func TestMinTimeToReach(t *testing.T) {
	for _, tc := range []struct {
		moveTime [][]int
		time     int
	}{
		{moveTime: [][]int{{0, 4}, {4, 4}}, time: 6},
		{moveTime: [][]int{{0, 0, 0}, {0, 0, 0}}, time: 3},
		{moveTime: [][]int{{0, 1}, {1, 2}}, time: 3},
	} {
		time := minTimeToReach(tc.moveTime)
		assert.Equal(t, tc.time, time)
	}
}
