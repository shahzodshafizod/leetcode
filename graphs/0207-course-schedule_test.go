package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCanFinish$
func TestCanFinish(t *testing.T) {
	for _, tc := range []struct {
		numCourses    int
		prerequisites [][]int
		can           bool
	}{
		{numCourses: 2, prerequisites: [][]int{{1, 0}}, can: true},
		{numCourses: 2, prerequisites: [][]int{{1, 0}, {0, 1}}, can: false},
		{numCourses: 6, prerequisites: [][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, can: true},
		{numCourses: 6, prerequisites: [][]int{{1, 0}, {2, 1}, {5, 2}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, can: false},
		{numCourses: 7, prerequisites: [][]int{{0, 3}, {1, 0}, {2, 1}, {4, 5}, {5, 6}, {6, 4}}, can: false},
		{numCourses: 0, prerequisites: [][]int{}, can: true},
	} {
		can := canFinish(tc.numCourses, tc.prerequisites)
		assert.Equal(t, tc.can, can)
	}
}
