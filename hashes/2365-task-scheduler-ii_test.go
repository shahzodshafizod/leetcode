package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestTaskSchedulerII$
func TestTaskSchedulerII(t *testing.T) {
	for _, tc := range []struct {
		tasks []int
		space int
		days  int64
	}{
		{tasks: []int{1, 2, 1, 2, 3, 1}, space: 3, days: 9},
		{tasks: []int{5, 8, 8, 5}, space: 2, days: 6},
	} {
		days := taskSchedulerII(tc.tasks, tc.space)
		assert.Equal(t, tc.days, days)
	}
}
