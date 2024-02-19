package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestKClosest$
func TestKClosest(t *testing.T) {
	for _, tc := range []struct {
		points        [][]int
		k             int
		closestPoints [][]int
	}{
		{points: [][]int{{1, 3}, {-2, 2}}, k: 1, closestPoints: [][]int{{-2, 2}}},
		{points: [][]int{{3, 3}, {5, -1}, {-2, 4}}, k: 2, closestPoints: [][]int{{3, 3}, {-2, 4}}},
	} {
		closestPoints := kClosest(tc.points, tc.k)
		assert.Equal(t, tc.closestPoints, closestPoints)
	}
}
