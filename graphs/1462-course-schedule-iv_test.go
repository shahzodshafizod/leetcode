package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestCheckIfPrerequisite$
func TestCheckIfPrerequisite(t *testing.T) {
	for _, tc := range []struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
		answer        []bool
	}{
		{numCourses: 2, prerequisites: [][]int{{1, 0}}, queries: [][]int{{0, 1}, {1, 0}}, answer: []bool{false, true}},
		{numCourses: 2, prerequisites: [][]int{}, queries: [][]int{{1, 0}, {0, 1}}, answer: []bool{false, false}},
		{numCourses: 3, prerequisites: [][]int{{1, 2}, {1, 0}, {2, 0}}, queries: [][]int{{1, 0}, {1, 2}}, answer: []bool{true, true}},
		{numCourses: 3, prerequisites: [][]int{{1, 0}, {2, 0}}, queries: [][]int{{0, 1}, {2, 0}}, answer: []bool{false, true}},
		{numCourses: 5, prerequisites: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, queries: [][]int{{0, 4}, {4, 0}, {1, 3}, {3, 0}}, answer: []bool{true, false, true, false}},
		{numCourses: 4, prerequisites: [][]int{{2, 3}, {2, 1}, {0, 3}, {0, 1}}, queries: [][]int{{0, 1}, {0, 3}, {2, 3}, {3, 0}, {2, 0}, {0, 2}}, answer: []bool{true, true, true, false, false, false}},
		{numCourses: 5, prerequisites: [][]int{{4, 3}, {4, 1}, {4, 0}, {3, 2}, {3, 1}, {3, 0}, {2, 1}, {2, 0}, {1, 0}}, queries: [][]int{{1, 4}, {4, 2}, {0, 1}, {4, 0}, {0, 2}, {1, 3}, {0, 1}}, answer: []bool{false, true, false, true, false, false, false}},
	} {
		answer := checkIfPrerequisite(tc.numCourses, tc.prerequisites, tc.queries)
		assert.Equal(t, tc.answer, answer)
	}
}
