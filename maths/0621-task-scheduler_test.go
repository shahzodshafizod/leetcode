package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestLeastInterval$
func TestLeastInterval(t *testing.T) {
	for _, tc := range []struct {
		tasks     []byte
		n         int
		intervals int
	}{
		{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 2, intervals: 8},
		{tasks: []byte{'A', 'C', 'A', 'B', 'D', 'B'}, n: 1, intervals: 6},
		{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 3, intervals: 10},
		{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C'}, n: 2, intervals: 8},
		{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'D', 'D', 'E'}, n: 2, intervals: 11},
	} {
		intervals := leastInterval(tc.tasks, tc.n)
		assert.Equal(t, tc.intervals, intervals)
	}
}
