package heaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./heaps/ -run ^TestMaxAverageRatio$
func TestMaxAverageRatio(t *testing.T) {
	for _, tc := range []struct {
		classes [][]int
		extraStudents int
		ratio float64
	}{
		{classes: [][]int{{1, 2}, {3, 5}, {2, 2}}, extraStudents: 2, ratio: 0.78333},
		{classes: [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, extraStudents: 4, ratio: 0.53485},
	} {
		ratio := maxAverageRatio(tc.classes, tc.extraStudents)
		assert.InDelta(t, tc.ratio, ratio, 0.00001)
	}
}
