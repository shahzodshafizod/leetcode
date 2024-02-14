package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestCountStudents$
func TestCountStudents(t *testing.T) {
	for _, tc := range []struct {
		students   []int
		sandwiches []int
		count      int
	}{
		{students: []int{1, 1, 0, 0}, sandwiches: []int{0, 1, 0, 1}, count: 0},
		{students: []int{1, 1, 1, 0, 0, 1}, sandwiches: []int{1, 0, 0, 0, 1, 1}, count: 3},
	} {
		count := countStudents(tc.students, tc.sandwiches)
		assert.Equal(t, tc.count, count)
	}
}
